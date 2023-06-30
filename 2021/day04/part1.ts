import { readFileSync } from 'fs';
import { chunk } from 'lodash';

const lines = readFileSync('day04/example.txt', 'utf-8').split('\n');

const drawNumbers = lines[0].split(',').map(n => Number.parseInt(n));

const cards = chunk(lines.slice(1),6).map(
    c => c.slice(1).join(' ').replaceAll(' ', '')
);

// console.log(cards);
