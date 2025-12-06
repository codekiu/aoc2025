import { log } from "console";
import fs from "fs";

function main(input: string) {
  const file = fs.readFileSync(input, "utf-8");

  const lines = file.split("\n");
  let result = 0;

  for (const line of lines) {
    let bankOfBatteries: number[] = [];
    const bankMaxSize = 12;
    const numbers = line.split("");
    let amountOfNumbersToCalculate = numbers.length;

    for (const numberStr of numbers) {
      const number = Number(numberStr);
      bankOfBatteries = process(
        bankOfBatteries,
        amountOfNumbersToCalculate,
        bankMaxSize,
        number,
      );
      amountOfNumbersToCalculate--;
    }
    result += getBatteryNumber(bankOfBatteries);
  }

  console.log("Result: ", result);
}

main("input");

function getBatteryNumber(batteries: Array<number>): number {
  return Number(batteries.join(""));
}

function process(
  bankOfBatteries: number[],
  amountOfNumbersToCalculate: number,
  bankMaxSize: number,
  currentBattery: number,
): number[] {
  let wasInserted = false;

  for (let i = 0; i < bankMaxSize; i++) {
    if (!bankOfBatteries[i]) {
      bankOfBatteries[i] = currentBattery;
      wasInserted = true;
      break;
    }
    if (amountOfNumbersToCalculate >= bankMaxSize - i) {
      if (bankOfBatteries[i] < currentBattery) {
        bankOfBatteries[i] = currentBattery;
        // clean up the batteries to the right
        for (let x = i + 1; x < bankMaxSize; x++) {
          bankOfBatteries[x] = 0;
        }
        wasInserted = true;
        break;
      }
    }
  }

  return bankOfBatteries;
}
