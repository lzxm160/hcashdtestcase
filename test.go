package main

import (
	"github.com/HcashOrg/hcashrpcclient"
	"github.com/HcashOrg/hcashutil"
	// "github.com/HcashOrg/hcashd/wire"
	"io/ioutil"
	// "log"
	"path/filepath"
	"time"
	"bufio"
	"github.com/HcashOrg/hcashd/chaincfg/chainhash"
	"github.com/HcashOrg/hcashwallet/loader"
)
func test1() {
	// Load the certificate for the TLS connection which is automatically
	// generated by hcashd when it starts the RPC server and doesn't already
	// have one.
	hcashdHomeDir := hcashutil.AppDataDir(".hcashd", false)
	certs, err := ioutil.ReadFile(filepath.Join(hcashdHomeDir, "rpc.cert"))
	if err != nil {
		log.Fatal(err)
	}

	// Create a new RPC client using websockets.  Since this example is
	// not long-lived, the connection will be closed as soon as the program
	// exits.
	connCfg := &hcashrpcclient.ConnConfig{
		Host:         "localhost:12101",
		Endpoint:     "ws",
		User:         "bitcoinrpc",
		Pass:         "123456",
		Certificates: certs,
	}
	client, err := hcashrpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Shutdown()

	// Query the RPC server for the genesis block using the "getblock"
	// command with the verbose flag set to true and the verboseTx flag
	// set to false.
	genesisHashStr := "000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f"
	blockHash, err := chainhash.NewHashFromStr(genesisHashStr)
	if err != nil {
		log.Fatal(err)
	}
	block, err := client.GetBlockVerbose(blockHash, false)
	if err != nil {
		log.Fatal(err)
	}

	// Display some details about the returned block.
	log.Printf("Hash: %v\n", block.Hash)
	log.Printf("Previous Block: %v\n", block.PreviousHash)
	log.Printf("Next Block: %v\n", block.NextHash)
	log.Printf("Merkle root: %v\n", block.MerkleRoot)
	log.Printf("Timestamp: %v\n", time.Unix(block.Time, 0).UTC())
	log.Printf("Confirmations: %v\n", block.Confirmations)
	log.Printf("Difficulty: %f\n", block.Difficulty)
	log.Printf("Size (in bytes): %v\n", block.Size)
	log.Printf("Num transactions: %v\n", len(block.Tx))
}
func test2() {
	hcashdHomeDir := hcashutil.AppDataDir(".hcashd", false)
	certs, err := ioutil.ReadFile(filepath.Join(hcashdHomeDir, "rpc.cert"))
	if err != nil {
		log.Fatal(err)
	}

	// Create a new RPC client using websockets.  Since this example is
	// not long-lived, the connection will be closed as soon as the program
	// exits.
	connCfg := &hcashrpcclient.ConnConfig{
		Host:         "localhost:12010",
		Endpoint:     "ws",
		User:         "bitcoinrpc",
		Pass:         "123456",
		Certificates: certs,
	}
	client, err := hcashrpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Shutdown()

	// Query the RPC server for the current block count and display it.
	blockCount, err := client.GetBlockCount()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Block count: %d", blockCount)
}
// networkDir returns the directory name of a network directory to hold wallet
// files.
func networkDir(dataDir string, chainParams *chaincfg.Params) string {
	netname := chainParams.Name

	// For now, we must always name the testnet data directory as "testnet"
	// and not "testnet" or any other version, as the chaincfg testnet
	// paramaters will likely be switched to being named "testnet" in the
	// future.  This is done to future proof that change, and an upgrade
	// plan to move the testnet data directory can be worked out later.
	switch chainParams.Net {
	case wire.TestNet2:
		netname = "testnet2"
	}

	return filepath.Join(dataDir, netname)
}
// createWallet prompts the user for information needed to generate a new wallet
// and generates the wallet accordingly.  The new wallet will reside at the
// provided path. The bool passed back gives whether or not the wallet was
// restored from seed, while the []byte passed is the private password required
// to do the initial sync.
func createWallet(cfg *config) error {
	dbDir := networkDir(cfg.AppDataDir.Value, activeNet.Params)
	stakeOptions := &loader.StakeOptions{
		VotingEnabled: cfg.EnableVoting,
		AddressReuse:  cfg.ReuseAddresses,
		TicketAddress: cfg.TicketAddress,
		TicketFee:     cfg.TicketFee.ToCoin(),
	}
	loader := loader.NewLoader(activeNet.Params, dbDir, stakeOptions,
		cfg.AddrIdxScanLen, cfg.AllowHighFees, cfg.RelayFee.ToCoin())

	reader := bufio.NewReader(os.Stdin)
	privPass, pubPass, seed, err := prompt.Setup(reader,
		[]byte(wallet.InsecurePubPassphrase), []byte(cfg.createPass), []byte(cfg.WalletPass))
	if err != nil {
		return err
	}

	fmt.Println("Creating the wallet...")
	_, err = loader.CreateNewWallet(pubPass, privPass, seed)
	if err != nil {
		return err
	}

	fmt.Println("The wallet has been created successfully.")

	return nil
}
func testcreatewallet() {
	tcfg, _, err := loadConfig()
	if err != nil {
		fmt.Println(err) 
	}
}
func main() {
	test2()
}