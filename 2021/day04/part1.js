"use strict";
exports.__esModule = true;
var fs_1 = require("fs");
var lodash_1 = require("lodash");
var lines = (0, fs_1.readFileSync)('day04/example.txt', 'utf-8').split('\n');
var drawNumbers = lines[0].split(',').map(function (n) { return Number.parseInt(n); });
var cards = (0, lodash_1.chunk)(lines.slice(1), 6).map(function (c) { return c.slice(1).join(' ').replaceAll(' ', ''); });
// console.log(cards);
