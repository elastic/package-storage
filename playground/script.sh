# Command to be used:

#promote mysql/1.3.2 snapshot staging


rm -rf build
mkdir -p build
cd build

# Check out both branches to be able to move the package from one to an other
git clone --single-branch --branch snapshot https://github.com/elastic/package-storage.git snapshot
git clone --single-branch --branch staging https://github.com/elastic/package-storage.git staging

# Make sure directory exists
mkdir -p staging/packages/snapshot
mv snapshot/packages/snapshot/0.0.1 staging/packages/snapshot/0.0.1

cd staging
git add "packages/snapshot/0.0.1"
git commit -a -m 'Adding package mysql/1.3.2'

mage build
docker build .

cd ../snapshot
git commit -a -m 'Remove package mysql/1.3.2'

mage build
docker build .

# If all went ok, push the commits to each branch

# Commit each change

# Run quick validation checks for both repo changes

# Push staging first, then snapshot
# Done