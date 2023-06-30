"use strict";
var __spreadArray = (this && this.__spreadArray) || function (to, from, pack) {
    if (pack || arguments.length === 2) for (var i = 0, l = from.length, ar; i < l; i++) {
        if (ar || !(i in from)) {
            if (!ar) ar = Array.prototype.slice.call(from, 0, i);
            ar[i] = from[i];
        }
    }
    return to.concat(ar || Array.prototype.slice.call(from));
};
exports.__esModule = true;
var fs_1 = require("fs");
var input = (0, fs_1.readFileSync)('day03/input.txt', 'utf-8');
var diags = input.split('\n');
var runCounts = function (diags) {
    var counts = [];
    diags.forEach(function (l) { return l.split('').forEach(function (d, i) {
        if (d == '1')
            counts[i] = (counts[i] || 0) + 1;
    }); });
    return counts;
};
var binaryStringToInt = function (s) {
    var n = 0;
    s.split('').forEach(function (d) {
        n = n << 1;
        if (d == '1')
            n++;
    });
    return n;
};
var oxyGenRating = __spreadArray([], diags, true);
var co2ScrubRating = __spreadArray([], diags, true);
var _loop_1 = function (pos) {
    var countsOxy = runCounts(oxyGenRating);
    var countsCo2 = runCounts(co2ScrubRating);
    var filterOnOxy = '0';
    if (oxyGenRating.length / 2 <= countsOxy[pos]) {
        filterOnOxy = '1';
    }
    var filterOnCo2 = '0';
    if (co2ScrubRating.length / 2 > countsCo2[pos]) {
        filterOnCo2 = '1';
    }
    if (oxyGenRating.length > 1) {
        oxyGenRating = oxyGenRating.filter(function (d) { return d[pos] == filterOnOxy; });
    }
    if (co2ScrubRating.length > 1) {
        co2ScrubRating = co2ScrubRating.filter(function (d) { return d[pos] == filterOnCo2; });
    }
};
for (var pos = 0; oxyGenRating.length > 1 || co2ScrubRating.length > 1; pos++) {
    _loop_1(pos);
}
console.log(oxyGenRating[0], co2ScrubRating[0]);
console.log(binaryStringToInt(oxyGenRating[0]), binaryStringToInt(co2ScrubRating[0]));
console.log(binaryStringToInt(oxyGenRating[0]) * binaryStringToInt(co2ScrubRating[0]));
