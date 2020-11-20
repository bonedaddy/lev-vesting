pragma solidity >=0.4.24 <=0.8.0;

import "./math/SafeMath.sol";
import "./interfaces/ERC20I.sol";

contract Vesting {

    uint256 public startBlock;
    uint256 public endBlock;

    ERC20Interface private levI; 
    
    constructor(address _levTokenAddress) {
        levI = ERC20Interface(_levTokenAddress);
    }

    function prepareVest(uint256 _amountToVest) public {
        require(levI.transferFrom(msg.sender, address(this), _amountToVest));
    }
}