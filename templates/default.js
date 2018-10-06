const chalk = require('chalk');

module.exports = ({hexString, colorName}) => {
  console.log('');
  console.log(chalk.bgHex(hexString)('       '));
  console.log(chalk.bgHex(hexString)('       ') + ' Your commit color is ' + colorName);
  console.log(chalk.bgHex(hexString)('       '));
  console.log('#' + hexString);
  console.log('');
}
