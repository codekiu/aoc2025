import { log } from "console";
import fs from "fs";

function main(input: string) {
  const file = fs.readFileSync(input, "utf-8");

  const lines = file.split("\n");
  const min = 1;
  const max = 9;
  let result = 0;

  lines.forEach((line) => {
    let batOne = 0;
    let batTwo = 0;
    const numbers = line.split("");
    let limit = numbers.length;

    numbers.forEach((numberStr) => {
      const number = Number(numberStr);

      if (batOne !== 0 && number > batTwo) {
        batTwo = number;
      }

      if (number > batOne && limit > 1) {
        batOne = number;
        batTwo = 0;
      }
      limit--;
    });
    const total = Number(batOne.toString() + batTwo.toString());
    result += total;
  });

  console.log("Result: ", result);
}

main("input");
