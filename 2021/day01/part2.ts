import { readFileSync } from 'fs';

const input = readFileSync('day01/input.txt', 'utf-8');
const depths = input.split('\n').map(i => Number.parseInt(i));

let countIncrements = 0;
for (let i = 1; i < depths.length-2; i++) {
    const windowA = depths[i-1] + depths[i] + depths[i+1];
    const windowB = depths[i] + depths[i+1] + depths[i+2];
    if (windowB > windowA) {
        countIncrements++;
    }
}

console.log(countIncrements);