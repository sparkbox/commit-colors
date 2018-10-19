#!/usr/bin/env node

const chalk = require('chalk');
const namer = require('color-namer');
const defaultTemplate = require('./templates/default');

const hexString = process.argv[2];
const colorName = namer(hexString).pantone[0].name;

let template = defaultTemplate;

if (process.argv[3]) {
  try {
    template = require(process.argv[3]);
  } catch(e) {
    console.error(`Couldn't find commit-color template: ${process.argv[3]}`);
  }
}

template({hexString, colorName});

