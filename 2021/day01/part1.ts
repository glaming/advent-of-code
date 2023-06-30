import { readFileSync } from 'fs';

const input = readFileSync('day01/input.txt', 'utf-8');
const depths = input.split('\n').map(i => Number.parseInt(i));

let countIncrements = 0;
for (let i = 1; i < depths.length; i++) {
    if (depths[i] > depths[i-1]) {
        countIncrements++;
    }
}

console.log(countIncrements);