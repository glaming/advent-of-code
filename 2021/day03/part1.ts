import { readFileSync } from 'fs';

const input = readFileSync('day03/input.txt', 'utf-8');
const diags = input.split('\n');

let counts: number[] = [];
diags.forEach(l => l.split('').forEach((d, i) => {
    if (d == '1') counts[i] = (counts[i] || 0) + 1;
}));

let gamma: number = 0, epsilon: number = 0;
counts.forEach((c, i) => {
    gamma = gamma << 1;
    epsilon = epsilon << 1

    if (diags.length / 2 < c) {
        gamma++
    } else {
        epsilon++
    }
});

console.log(gamma, epsilon, gamma * epsilon);