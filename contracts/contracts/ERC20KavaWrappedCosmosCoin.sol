// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

/// @title ERC20 Token Contract for 0g-chain Wrapped Cosmos Coin
/// @notice Tokens are backed one-for-one by cosmos-sdk coins held in the module account.
/// @dev This contract is owned and deployed by the evmutil module of 0g-chain.
/// @custom:security-contact security@0g.ai
contract ERC200gChainWrappedCosmosCoin is ERC20, Ownable {
    uint8 private immutable _decimals;

    /// @notice Constructor to initialize the ERC20 token with mint and burn permissions.
    /// @param name The name of the ERC20 token.
    /// @param symbol The symbol of the ERC20 token.
    /// @param decimals_ The number of decimals of the ERC20 token.
    constructor(
        string memory name,
        string memory symbol,
        uint8 decimals_
    ) ERC20(name, symbol) {
        _decimals = decimals_;
    }

    /// @notice Returns the number of decimal places for the token.
    /// @return The number of decimal places.
    function decimals() public view override returns (uint8) {
        return _decimals;
    }

    /// @notice Mints new tokens to a specified address.
    /// @dev Can only be called by the contract owner.
    /// @param to The address to which new tokens are minted.
    /// @param amount The number of tokens to mint.
    function mint(address to, uint256 amount) external onlyOwner {
        _mint(to, amount);
    }

    /// @notice Burns tokens from a specified address.
    /// @dev Can only be called by the contract owner.
    /// @param from The address from which tokens are burned.
    /// @param amount The number of tokens to burn.
    function burn(address from, uint256 amount) external onlyOwner {
        _burn(from, amount);
    }
}
