"use strict";
exports.__esModule = true;
var fs_1 = require("fs");
var input = (0, fs_1.readFileSync)('day01/input.txt', 'utf-8');
var depths = input.split('\n').map(function (i) { return Number.parseInt(i); });
var countIncrements = 0;
for (var i = 1; i < depths.length - 2; i++) {
    var windowA = depths[i - 1] + depths[i] + depths[i + 1];
    var windowB = depths[i] + depths[i + 1] + depths[i + 2];
    if (windowB > windowA) {
        countIncrements++;
    }
}
console.log(countIncrements);
