import { readFileSync } from 'fs';

const input = readFileSync('day02/input.txt', 'utf-8');
const instructions = input.split('\n').map(l => {
    const s = l.split(' ');
    return {direction: s[0], value: Number.parseInt(s[1])};
});


let pos = {depth: 0, horizontal: 0};

instructions.forEach(i => {
    switch (i.direction) {
        case 'up':
            pos.depth = pos.depth - i.value;
            break
        case 'down':
            pos.depth = pos.depth + i.value;
            break
        case 'forward':
            pos.horizontal = pos.horizontal + i.value;
            break
    }
});

console.log(pos);