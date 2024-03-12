#!/bin/bash

echo "Please enter Leetcode number:"
read problemNumber

echo "Please enter function name:"
read funcName

echo "Please enter parameters for the function (separated by spaces):"
read params

echo "Please enter return type:"
read returnType

echo "export default function $funcName($params): $returnType {
  // your code here
}" >"./solutions/$problemNumber.ts"

echo "import { expect } from 'chai';
import $funcName from '@solutions/$problemNumber';

describe('Testing LC $problemNumber', () => {
  it('test case 1', () => {
    const arg1 = 0;
    const result = $funcName(arg1);
    const expected = 0;
    expect(result).to.equal(expected);
});
});" >"./test/$problemNumber.test.ts"

echo "$problemNumber.ts has been created in solutions folder"
echo "$problemNumber.test.ts has been created in test folder"
