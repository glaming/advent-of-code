"use strict";
exports.__esModule = true;
var fs_1 = require("fs");
var input = (0, fs_1.readFileSync)('day02/input.txt', 'utf-8');
var instructions = input.split('\n').map(function (l) {
    var s = l.split(' ');
    return { direction: s[0], value: Number.parseInt(s[1]) };
});
var pos = { depth: 0, horizontal: 0 };
instructions.forEach(function (i) {
    switch (i.direction) {
        case 'up':
            pos.depth = pos.depth - i.value;
            break;
        case 'down':
            pos.depth = pos.depth + i.value;
            break;
        case 'forward':
            pos.horizontal = pos.horizontal + i.value;
            break;
    }
});
console.log(pos);
