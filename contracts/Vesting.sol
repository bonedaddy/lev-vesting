pragma solidity >=0.4.24 <=0.8.0;

import "./math/SafeMath.sol";
import "./interfaces/ERC20I.sol";
import "./interfaces/DateTimeInterface.sol";

contract Vesting {

    uint256 public startTime;
    uint256 public endTime;
    uint256 public currentRelease;
    uint256 public releaseAmount;
    address public receiver;

    ERC20Interface private levI; 
    DateTimeInterface private dateI;

    struct Release {
        uint256 timestamp;
        uint256 released;
    }

    mapping (uint256 => Release) public releases;

    constructor(address _levTokenAddress, address _dateTimeContract) {
        levI = ERC20Interface(_levTokenAddress);
        dateI = DateTimeInterface(_dateTimeContract);
    }

    function prepareVest(uint256 _amountToVest, address _receiver) public {
        require(levI.transferFrom(msg.sender, address(this), _amountToVest));
        startTime = dateI._now();
        endTime = dateI.addMonths(startTime, 12);
        releaseAmount = _amountToVest / 12;
        receiver = _receiver;
        currentRelease = 1;
        for (uint i = 1; i <= 12; i++) {
            releases[i] = Release({
                timestamp: dateI.addMonths(startTime, i),
                released: 0
            });
        }
    }

    function release() public {
        require(isPrepared() == true);
        uint256 timestamp = dateI._now();
        require(releases[currentRelease].released == 0, "already released");
        releases[currentRelease].released = 1;
        require(timestamp >= releases[currentRelease].timestamp, "release timestamp not yet passed");
        require(levI.transfer(receiver, releaseAmount));
        currentRelease += 1;
    }

    function isReleased(uint256 _index) public view returns (bool) {
        bool released = false;
        if (releases[_index].released == 1) {
            released = true;
        }
        return released;
    }

    function isPrepared() public view returns (bool) {
        bool prepared = false;
        if (receiver != address(0) && releaseAmount > 0 && currentRelease > 0) {
            prepared = true;
        }
        return prepared;
    }
}