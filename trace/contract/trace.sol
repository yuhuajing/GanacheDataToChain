// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.8.2 <0.9.0;

contract AgriTra {
    mapping(string => OrgInfo) public tracedata;

    struct OrgInfo {
        string number;
        string workamount;
        string persion;
        string workmethod;
        string worktime;
        string remarks;
    }

    constructor() {}

    function addorupdateData(
        string memory number,
        string memory workamount,
        string memory persion,
        string memory workmethod,
        string memory worktime,
        string memory remarks
    ) external {
        tracedata[number].workamount = workamount;
        tracedata[number].persion = persion;
        tracedata[number].workmethod = workmethod;
        tracedata[number].worktime = worktime;
        tracedata[number].remarks = remarks;
    }

    function tracedataByNum(string memory number)
    external
    view
    returns (OrgInfo memory)
    {
        return tracedata[number];
    }
}
