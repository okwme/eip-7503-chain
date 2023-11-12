// to deploy local
//npx hardhat node
//npx hardhat run --network localhost scripts/deploy.js
const hre = require("hardhat");
const path = require("node:path");

const fs = require("fs").promises;

const { deployContracts } = require("./utils");

async function writedata(path, data) {
  // await fs.writeFile(path, data, function (err, result) {
  //   if (err) console.log('error', err);
  // })
  try {
    names = await fs.writeFile(path, data);
  } catch (e) {
    console.log("e", e);
  }
}

async function copyContractABI(a, b) {
  try {
    write = await fs.copyFile(a, b);
  } catch (e) {
    console.log("e", e);
  }
}

async function copyABI(name) {
  var networkinfo = await hre.ethers.provider.getNetwork();
  console.log({ networkinfo })
  console.log(`--copy ${name} ABI`);
  var pathname = path.join(
    __dirname,
    "..",
    "artifacts",
    "contracts",
    `${name}.sol`,
    `${name}.json`
  );
  var copy = path.join(
    __dirname,
    "..",
    "contractMetadata",
    "ABI-" + String(networkinfo["name"]) + `-${name}.json`
  );
  await copyContractABI(pathname, copy);
  console.log("-- OK");
}

async function saveAddress(contract, name) {
  console.log("-save json for " + name);
  var networkinfo = await hre.ethers.provider.getNetwork();
  var newAddress = await contract.address;
  var savePath = path.join(
    __dirname,
    "..",
    "contractMetadata",
    String(networkinfo["name"]) + "-" + String(name) + ".json"
  );
  var objToWrite = {
    address: newAddress,
    chain: networkinfo,
  };
  await writedata(savePath, JSON.stringify(objToWrite));
}

async function main() {
  console.log("Deploy to chain:");
  const network = await hre.ethers.provider.getNetwork()
  console.log({ network })
  const { mint } = await deployContracts();
  await copyABI("Mint");
  await saveAddress(mint, "Mint");
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });

module.exports = { deployContracts };
