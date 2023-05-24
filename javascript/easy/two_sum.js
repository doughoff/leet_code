/**
 * twoSum
 * @param { Array < number >} arr
 * @param { number; } target
 * 
 * @returns { Array < number >}
 */
function twoSum(arr, target) {
  const result = [];

  for (let i = 0; i < arr.length; i++) {
    const vInner = arr[i];
    for (let j = i + 1; j < arr.length; j++) {
      const vOuter = arr[j];
      if (vInner + vOuter === target) {
        result.push(i, j + i);
        return result;
      }
    }
  }

  return result;
}

/**
 * twoSum
 * @param { Array < number >} arr
 * @param { number; } target
 * 
 * @returns { Array < number >}
 */
function twoSum2(arr, target) {
  const map = {};

  for (let i = 0; i < arr.length; i++) {
    const v = arr[i];
    const diff = target - v;

    if (diff in map) {
      return [map[diff], i];
    }

    map[v] = i;
  }

  return [];
}

function main() {
  console.log(twoSum([2, 7, 11, 15], 9), [0, 1]);

  const start1 = new Date().getTime();
  const bigArr = [];

  for (let i = 0; i < 100_000_000; i++) {
    // get random number 1 - 1_000_000
    const rand = Math.floor(Math.random() * 1_000_000) + 1;
    bigArr.push(rand);
  }

  console.log(bigArr.length);
  const end1 = new Date().getTime();
  console.log(`Time: ${end1 - start1}ms`);


  const target = 1_500_000;

  // get Start timer
  const start = new Date().getTime();
  const [a, b] = twoSum2(bigArr, target);

  console.log(a, b, 'values', bigArr[a], bigArr[b], 'sum', bigArr[a] + bigArr[b], target);
  // get End timer
  const end = new Date().getTime();
  console.log(`Time: ${end - start}ms`);
}

main();