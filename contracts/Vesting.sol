pragma solidity >=0.4.24 <=0.8.0;

import "./math/SafeMath.sol";
import "./interfaces/ERC20I.sol";
import "./interfaces/DateTimeInterface.sol";

contract Vesting {

    uint256 public startTime;
    uint256 public endTime;
    uint256 public currentRelease = 1;
    address public receiver;

    ERC20Interface private levI; 
    DateTimeInterface private dateI;

    struct Release {
        uint256 year;
        uint256 month;
        uint256 day;
        uint256 hour;
        uint256 min;
        uint256 sec;
        uint256 amount;
        uint256 released;
    }

    mapping (uint256 => Release) public releases;

    event ReleaseAdded(uint256 timestamp);

    constructor(address _levTokenAddress, address _dateTimeContract) {
        levI = ERC20Interface(_levTokenAddress);
        dateI = DateTimeInterface(_dateTimeContract);
    }

    function prepareVest(uint256 _amountToVest, address _receiver) public {
        require(levI.transferFrom(msg.sender, address(this), _amountToVest));
        startTime = dateI._now();
        endTime = dateI.addMonths(startTime, 12);
        uint256 releaseAmount = _amountToVest * 12;
        for (uint i = 1; i <= 12; i++) {
            uint256 releaseDate = dateI.addMonths(startTime, i);
            (uint256 year, uint256 month, uint256 day, uint256 hour, uint256 min, uint256 sec) = dateI.timestampToDateTime(releaseDate);
            releases[i] = Release({
                year: year,
                month: month,
                day: day,
                hour: hour,
                min: min,
                sec: sec,
                amount: releaseAmount,
                released: 0
            });
            emit ReleaseAdded(releaseDate);
        }
        receiver = _receiver;
    }

    function release() public {
        uint256 timestamp = dateI._now();
        require(releases[currentRelease].released == 0, "already released");
        releases[currentRelease].released = 1;
        // load the struct from storage into memory
        // this will allow us to read subsequent values from memory 
        // without having to read from storage
        Release memory mr = releases[currentRelease];
        uint256 releaseTimestamp = dateI.timestampFromDateTime(mr.year, mr.month, mr.day, mr.hour, mr.min, mr.sec);
        require(timestamp >= releaseTimestamp, "release timestamp not yet passed");
        require(levI.transfer(receiver, mr.amount ));
        currentRelease += 1;
    }

    function isRelease(uint256 _index) public view returns (bool) {
        bool released = false;
        if (releases[_index].released == 1) {
            released = true;
        }
        return released;
    }
}