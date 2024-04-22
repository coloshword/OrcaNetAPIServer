echo "Building the project..."
make all
echo "Launching OrcaNet to initialize..."
./OrcaNet/OrcaNet &
ORCANET_PID=$!

# Wait for 1 second before killing OrcaNet
sleep 1
echo "Stopping OrcaNet..."
kill $ORCANET_PID
sleep 2
clear
echo "Checking if wallet exists, will prompt creation if not created"
./OrcaWallet/btcwallet --create
echo "Build complete"
