import { readFileSync } from 'fs';

const input = readFileSync('day03/input.txt', 'utf-8');
const diags = input.split('\n');

const runCounts = (diags: string[]) => {
    let counts: number[] = [];

    diags.forEach(l => l.split('').forEach((d, i) => {
        if (d == '1') counts[i] = (counts[i] || 0) + 1;
    }));

    return counts;
}

const binaryStringToInt = (s: string) => {
    let n: number = 0;
    s.split('').forEach(d => {
        n = n << 1
        if (d == '1') n++;
    });
    return n;
}

let oxyGenRating = [...diags];
let co2ScrubRating = [...diags];
for (let pos = 0; oxyGenRating.length > 1 || co2ScrubRating.length > 1 ; pos++) {

    const countsOxy = runCounts(oxyGenRating);
    const countsCo2 = runCounts(co2ScrubRating);

    let filterOnOxy = '0';
    if (oxyGenRating.length / 2 <= countsOxy[pos]) {
        filterOnOxy = '1';
    }

    let filterOnCo2 = '0';
    if (co2ScrubRating.length / 2 > countsCo2[pos]) {
        filterOnCo2 = '1';
    }

    if (oxyGenRating.length > 1) {
        oxyGenRating = oxyGenRating.filter(d => d[pos] == filterOnOxy);
    }
    if (co2ScrubRating.length > 1) {
        co2ScrubRating = co2ScrubRating.filter(d => d[pos] == filterOnCo2);
    }
}

console.log(oxyGenRating[0], co2ScrubRating[0]);
console.log(binaryStringToInt(oxyGenRating[0]), binaryStringToInt(co2ScrubRating[0]));
console.log(binaryStringToInt(oxyGenRating[0]) * binaryStringToInt(co2ScrubRating[0]));
