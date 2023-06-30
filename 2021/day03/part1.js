"use strict";
exports.__esModule = true;
var fs_1 = require("fs");
var input = (0, fs_1.readFileSync)('day03/input.txt', 'utf-8');
var diags = input.split('\n');
var counts = [];
diags.forEach(function (l) { return l.split('').forEach(function (d, i) {
    if (d == '1')
        counts[i] = (counts[i] || 0) + 1;
}); });
console.log(counts);
var gamma = 0, epsilon = 0;
counts.forEach(function (c, i) {
    gamma = gamma << 1;
    epsilon = epsilon << 1;
    if (diags.length / 2 < c) {
        gamma++;
    }
    else {
        epsilon++;
    }
});
console.log(gamma, epsilon, gamma * epsilon);
