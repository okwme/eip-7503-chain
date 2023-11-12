.PHONY = tau circuit all2 choose_option option1 option2

tau_create:
	@echo "after creating the power of tau file place it in build/power_of_tau directory and update the make file";\
	snarkjs powersoftau new bn128 20 pot20_0000.ptau -v
	snarkjs powersoftau contribute pot20_0000.ptau pot20_0001.ptau --name="First contribution" --entropy=1234 -v
	snarkjs powersoftau prepare phase2 pot20_0001.ptau pot20_final.ptau -v


tau:
	@echo "check the snarkjs repo to download and after download place the file in build/power_of_tau directory";\

compile:
	@read -p "Enter the circuit name: " value; echo "$$value" > temp.txt;
	@echo "Please choose prover type:"
	@echo "1) c compiled prover"
	@echo "2) snarkjs wasm prover"
	@read -p "Enter number: " number; \
	case $$number in \
	1) $(MAKE) compile_c ;; \
	2) $(MAKE) compile_snarkjs ;; \
	*) echo "Invalid option" ;; \
	esac

compile_c:
	@circuit_name=`cat temp.txt`;\
	mkdir -p build/circuits/$$circuit_name;\
	cd build/circuits/$$circuit_name;\
	circom ../../../circuit/$$circuit_name.circom --r1cs --wasm --sym --c;\
	cd $$circuit_name_cpp && make && cd ..;\
	snarkjs groth16 setup $$circuit_name.r1cs ../../power_of_tau/powersOfTau28_hez_final_23.ptau $$circuit_name_0000.zkey;\
	snarkjs zkey contribute $$circuit_name_0000.zkey $$circuit_name_0001.zkey --entropy=1234 --name="1st Contributor Name" -v;\
	cd ../../..;\
	rm temp.txt
compile_snarkjs:
	@circuit_name=`cat temp.txt`;\
	mkdir -p build/circuits/$$circuit_name;\
	cd build/circuits/$$circuit_name;\
	circom ../../../circuit/$$circuit_name.circom --r1cs --wasm --sym;\
	snarkjs groth16 setup $$circuit_name.r1cs ../../power_of_tau/powersOfTau28_hez_final_23.ptau 0000.zkey;\
	snarkjs zkey contribute 0000.zkey 0001.zkey --entropy=1234 --name="1st Contributor Name" -v;\
	snarkjs zkey export verificationkey 0001.zkey verification_key.json;\
	cd ../../..;\
	rm temp.txt;\
	echo "Compiled. check build/circuits/$$circuit_name"


input_gen:
	@read -p "Enter the circuit name: " value; \
	python3.10 scripts/input_gen/$$value.py > build/circuits/$$value/input.json;\
	echo "Input generated at build/circuits/$$value/input.json"

prove:
	@read -p "Enter the circuit name: " value; \
	cd build/circuits/$$value;\
	snarkjs wtns calculate $$value\_js/$$value.wasm input.json witness.wtns;\
	snarkjs groth16 prove 0001.zkey witness.wtns proof.json public.json;\
	echo "Proof generated at build/circuits/$$value/proof.json"

verify:
	@read -p "Enter the circuit name: " value; \
	cd build/circuits/$$value;\
	snarkjs groth16 verify verification_key.json public.json proof.json;\

clean:
	@read -p "Enter the circuit name: " value; \
	rm -rf build/circuits/$$value;\
	echo "Cleaned build/circuits/$$value"
