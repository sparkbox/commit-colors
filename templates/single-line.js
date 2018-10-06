#!/usr/bin/env node

const chalk = require('chalk');

const getTextColor = (hexString) => {
  const decimal = Buffer.from(hexString, 'hex');
  const r = 255 - decimal[0];
  const g = 255 - decimal[1];
  const b = 255 - decimal[2];

  return Buffer.from([r, g, b]).toString('hex');
}


module.exports = ({ hexString, colorName }) => {
  const fgColor = getTextColor(hexString);
  console.log(chalk.hex(fgColor).bgHex(hexString)(` Your commit color is ${colorName} `), `#${hexString}`);
}
