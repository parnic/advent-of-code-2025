import * as fs from "node:fs";

const data = fs.readFileSync("inputs/02p.txt", "utf8");
const ranges = data.split(",");
const r = /^([0-9]+)\1$/gm;
const r2 = /^([0-9]+)\1+$/gm;
let sum = 0;
let sum2 = 0;
for (const range of ranges) {
  const vals = range.split("-");
  const min = Number(vals[0]);
  const max = Number(vals[1]);
  for (let i = min; i <= max; i++) {
    if (i.toString().match(r)) {
      sum += i;
    }
    if (i.toString().match(r2)) {
      sum2 += i;
    }
  }
}
console.log(sum);
console.log(sum2);
