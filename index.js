#!/usr/bin/env node

const chalk = require('chalk');
const namer = require('color-namer');

const hexString = process.argv[2];
const colorName = namer(hexString).pantone[0].name;

console.log('');
console.log(chalk.bgHex(hexString)('       '));
console.log(chalk.bgHex(hexString)('       ') + ' Your commit color is ' + colorName);
console.log(chalk.bgHex(hexString)('       '));
console.log('#' + hexString);
console.log('');
