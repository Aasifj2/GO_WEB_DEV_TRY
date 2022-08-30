package main

//CURRENT ACTIVE
import (
	//"crypto/elliptic/internal/nistec"
	"crypto/rand"
	"io/ioutil"
	"math"
	"strings"

	//"encoding"
	"encoding/hex"
	"encoding/json"
	"fmt"

	//"ioutil"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/multiformats/go-multiaddr"
	"gopkg.in/dedis/kyber.v2"
	"gopkg.in/dedis/kyber.v2/util/encoding"
	//"go.uber.org/zap/internal/exit"
)

func Key_Refresh(T int64, Peer_Count int64, peer_number string, peer_details_list []string, protocolID protocol.ID) {
	fmt.Println("****************Starting Key Refresh*********")
	// //Setting Constatnt term of the coeff as zero
	// Set_secret(curve.Scalar().Zero())

	// poly := []kyber.Scalar{}  // to store coefficients
	// share := []kyber.Scalar{} // to store share
	// alphas := []kyber.Point{} // to store alphas

	// var i int64

	// for i = 0; i < T; i++ {
	// 	poly = append(poly, curve.Scalar().Zero())
	// }

	// for i = 0; i < T; i++ {
	// 	alphas = append(alphas, curve.Point().Null())
	// }

	// for i = 0; i <= int64(Peer_Count); i++ {
	// 	share = append(share, curve.Scalar().Zero())
	// }

	// // to generate coefficients of the polynomial
	// Generate_Polynomial_coefficients(T, poly, peer_number)
	// // Generating the shares and storing in share array
	// Generate_share(int64(Peer_Count), T, poly, share, peer_number)
	// //Generating Alphas
	// Generate_Alphas(T, alphas, poly, peer_number)
	// //Adding Key Refresh Coeff + coeff till last refresh and storing it as Curr_coeff
	// //Updating Coeff[i]
	// Path := "vss/" + strconv.Itoa(my_index)
	// var val kyber.Scalar = curve.Scalar().Zero()
	// for i = 0; i < T; i++ {
	// 	path1 := Path + "/Curr_coeff" + strconv.Itoa(int(i)) + ".txt"
	// 	file, e2 := os.Open(path1)
	// 	if e2 != nil {
	// 		panic(e2)
	// 	}
	// 	val, _ = encoding.ReadHexScalar(curve, file)
	// 	val.Add(val, poly[i])
	// 	file.Close()
	// 	//Writing into the file
	// 	file, e2 = os.OpenFile(path1, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	// 	if e2 != nil {
	// 		panic(e2)
	// 	}
	// 	encoding.WriteHexScalar(curve, file, val)
	// 	file.Close()
	// }
	// //Updating Share[i]
	// for i = 0; i < T; i++ {
	// 	path1 := Path + "/Curr_Share" + strconv.Itoa(int(i)) + ".txt"
	// 	file, e2 := os.Open(path1)
	// 	if e2 != nil {
	// 		panic(e2)
	// 	}
	// 	val, _ = encoding.ReadHexScalar(curve, file)
	// 	val.Add(val, share[i])
	// 	file.Close()
	// 	//Writing into the file
	// 	file, e2 = os.OpenFile(path1, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	// 	if e2 != nil {
	// 		panic(e2)
	// 	}
	// 	encoding.WriteHexScalar(curve, file, val)
	// 	file.Close()
	// }
	// //Broadcasting alphas
	// fmt.Println("Brodcasting Alphas")
	// status_struct.Phase = 6
	// for i = 0; i < T; i++ {

	// 	send_data(peer_details_list, alphas[i].String(), fmt.Sprint(i), protocolID)

	//this
	// }
	// //Waiting for other peers to get to same phase
	// wait_until(6)

	// //Receiving alphas from other peers
	// Recieve_Alphas(Peer_Count, peer_number, T)
	// //Encrypting Shares and Broadcasting

	// fmt.Println("ENCRYPTING SHARES & Broadcasting")
	// //Defining Elgamal Curve
	// elg_curve := Setup()
	// //Reading Sender's elgamal Public key
	// path1 := "Data/" + peer_number
	// Sender_EPK_file, _ := ioutil.ReadFile(path1 + "/EPK.txt")
	// Sender_EPK_file, _ = hex.DecodeString(string(Sender_EPK_file))
	// Sender_EPK, _ := elg_curve.Point.FromAffineCompressed(Sender_EPK_file)
	// //Reading Sender's elgamal Secret Key
	// Sender_ESK_file, _ := ioutil.ReadFile(path1 + "/ESK.txt")
	// Sender_ESK_file, _ = hex.DecodeString(string(Sender_ESK_file))
	// Sender_ESK, _ := elg_curve.Scalar.SetBytes(Sender_ESK_file)

	// //Path to vss generated parameters
	// path := "vss/" + peer_number
	// status_struct.Phase = 7
	// for i = 0; i <= int64(Peer_Count); i++ {
	// 	if i == int64(my_index) {
	// 		continue
	// 	}
	// 	//reading epk from broadcast to encrypt it (for ith user)
	// 	_f, err := os.Open(path + "/Indivisual_Share" + strconv.Itoa(int(i)) + ".txt")
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	//share for ith user generated by current peer
	// 	share, _ := encoding.ReadHexScalar(curve, _f)

	// 	//Reading Elgamal Public key of ith user
	// 	data, err := os.ReadFile("Broadcast/" + strconv.Itoa(int(i)) + "/EPK.txt")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	data, _ = hex.DecodeString(string(data))
	// 	elg_curve := Setup()
	// 	EPK_receiver, _ := elg_curve.Point.FromAffineCompressed(data)

	// 	//Share to encrypt(in string format)
	// 	toEncrypt := share.String()

	// 	//Ecryption using( current peer's Secret key,current peer's public key, ith users(receivers) public key)
	// 	C1, C2, C3, _ := AuthEncryption(toEncrypt, Sender_ESK, Sender_EPK, EPK_receiver)

	// 	//Broadcasting the encrypted values.
	// 	var j = 1
	// 	for j = 1; j <= 3; j++ {
	// 		if j == 1 {

	// 			send_data(peer_details_list, fmt.Sprintf("%x", C1.ToAffineCompressed()), fmt.Sprint(i)+","+fmt.Sprint(j), protocolID)
	// 		}
	// 		if j == 2 {
	// 			send_data(peer_details_list, string(C2), fmt.Sprint(i)+","+fmt.Sprint(j), protocolID)

	// 		}
	// 		if j == 3 {
	// 			send_data(peer_details_list, hex.EncodeToString(C3), fmt.Sprint(i)+","+fmt.Sprint(j), protocolID)

	// 		}
	// 	}

	// }
	// wait_until(7)
	// //
	// fmt.Println("Decrypting Shares")
	// //Setting up elgamal curve
	// elg_curve = Setup()

	// //Reading elgamal Public key of current peer
	// path = "Data/" + peer_number

	// Reciever_EPK, _ := Get_EPK(path + "/EPK.txt")

	// Reciever_ESK, _ := Get_ESK(path + "/ESK.txt")

	// for i = 0; i <= int64(Peer_Count); i++ {
	// 	if i == int64(my_index) {
	// 		continue
	// 	}

	// 	//Reading Elgamal Public key of Sender(i)
	// 	path2 := "Broadcast/" + fmt.Sprint(i) + "/EPK.txt" //Get the epk of sender

	// 	Sender_EPK, _ := Get_EPK(path2)

	// 	//Reading the Encrypted data sent by ith user to current user from broadcast folder(channel)
	// 	C1_Data, _ := ioutil.ReadFile("Broadcast/" + fmt.Sprint(i) + "/Shares/To" + peer_number + "/C1.txt")
	// 	C2_Data, _ := ioutil.ReadFile("Broadcast/" + fmt.Sprint(i) + "/Shares/To" + peer_number + "/C2.txt")
	// 	C3_Data, _ := ioutil.ReadFile("Broadcast/" + fmt.Sprint(i) + "/Shares/To" + peer_number + "/C3.txt")

	// 	//Changing the data read into C1,C2,C3 format
	// 	data, _ := hex.DecodeString(string(C1_Data))
	// 	C1, _ := elg_curve.Point.FromAffineCompressed(data)
	// 	C2 := C2_Data
	// 	C3, _ := hex.DecodeString(string(C3_Data))
	// 	// C3 := C3_Data

	// 	//Decryption of shares(C1,C2,C3)
	// 	share, err := AuthDecryption(C1, C2, C3, Sender_EPK, Reciever_EPK, Reciever_ESK)

	// 	if err != nil {
	// 		fmt.Println("Error Decrypting")
	// 	}

	// 	//Saving the decrypted message into the received folder of current user
	// 	path2 = "Received/" + peer_number + "/Shares/share" + fmt.Sprint(i) + ".txt"
	// 	os.MkdirAll("Received/"+peer_number+"/Shares/", os.ModePerm)
	// 	file, _ := os.OpenFile(path2, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	// 	_, _ = fmt.Fprint(file, string(share))

	// }
	// fmt.Println(peer_number, "Verifying Shares")

	// G := Verify_Share(peer_number, int64(Peer_Count), int64(T))
	// fmt.Println(G.String())
	// path = "Received/Signing/" + peer_number
	// file, _ := os.Open(path + "/G.txt")
	// val1, _ := encoding.ReadHexScalar(curve, file)
	// file.Close()
	// G.Add(G, val1)
	// file, _ = os.Create(path)
	// encoding.WriteHexScalar(curve, file, G)
	// file.Close()
	fmt.Print("")
}

func check(peer_number string, f_i kyber.Scalar, T int64, alphas []kyber.Point) bool {
	//i int64, f_i kyber.Scalar, T int64, alphas []kyber.Point
	var v1 kyber.Point
	v1 = curve.Point().Null()
	// share_i := Readshare(i)
	i, _ := strconv.Atoi(peer_number)
	share_i := f_i
	v1.Mul(share_i, g)
	var j int64
	var sum kyber.Point
	sum = curve.Point().Null()
	for j = 0; j < T; j++ {
		X := math.Pow(float64(i), float64(j))
		x := curve.Scalar().SetInt64(int64(X))
		var v4 kyber.Point = curve.Point().Null()
		// v4 = Readalpha(j)
		v4 = alphas[j]
		var val kyber.Point
		val = curve.Point().Null()
		val.Mul(x, v4)
		sum.Add(val, sum)
	}
	fmt.Println(v1)
	fmt.Println(sum)
	if v1.Equal(sum) {
		fmt.Print("HURRAY:")
	} else {
		fmt.Println("SADDD")
	}
	return v1.Equal(sum)
}

func verify_GK(Peer_Count int64, T int64) bool {
	GK := Get_Group_Key(Peer_Count)
	fmt.Println("GROUP KEY IN VERIFY GK:", GK.String(), "\n\n")
	sum := curve.Scalar().Zero()
	var i int64
	for i = 1; i <= Peer_Count; i++ {
		path := "Received/" + strconv.Itoa(int(i)) + "/G.txt"
		file, _ := os.Open(path)
		temp, _ := encoding.ReadHexScalar(curve, file)
		lambda := Lambda(T, i)
		prod := curve.Scalar().Mul(lambda, temp)
		sum = sum.Add(sum, prod)
	}
	final := curve.Point().Mul(sum, g)
	if GK.Equal(final) {
		return true
	} else {
		return false
	}
}

//Rework to common message sending function
func keygen() {
	// if keygenFlag {
	// 	return
	// }
	// keygenFlag = true
	// peer_details_list := peer_details_list

	//current_flag = "1"
	status_struct.Phase = 1
	var protocolID protocol.ID = "/keygen/0.0.1"
	//Start Listener
	keygen_Stream_listener(p2p.Host)
	//Start Acknowledger
	host_acknowledge(p2p.Host)

	//Generate broadcast wait time
	time.Sleep(time.Second * 5)
	// fmt.Println(GeneratePrime(1024))

	peer_number := fmt.Sprint(my_index + 1)
	Peer_Count := len(peer_details_list)
	fmt.Println("PEERCOUNT:", Peer_Count)
	fmt.Println("MYINDEX:", peer_number)
	var T int64 = int64(Threshold)
	fmt.Println("THRESHOLD:", T)
	fmt.Println(peer_details_list)
	//./zebpay -f first phase
	//function
	// 1. first phase - Setup
	// 2. second phase - Keygen

	//First - they generate and store elgamal keys
	//SECOND - RUN p2p key is broadcasted --path for key
	//THIRD - Generation for shares --path
	//FOURTH - RUN P2P broadcasting the shares & presigning

	// fmt.Println(peer_details_list)
	// var choice int
	// choice = 1

	// fmt.Println("MY INDEX:--->", my_index)
	// fmt.Println("1. Key Setup")
	// fmt.Println("2. BroadCast Public Key")
	// fmt.Println("3. Comit ")
	// fmt.Println("4. DeComit")
	// fmt.Println("5  Generate Shares")
	// fmt.Println("6  Encrypt Shares and Broadcast")
	// fmt.Println("7  Decrypt Shares")
	// fmt.Println("8  Verify Shares")
	// fmt.Println("9  Signing Setup")
	// fmt.Println("10 Signing Commit")
	// fmt.Println("11 Decomit")
	// fmt.Println("12 Sign VSS Setup and Processing")
	// fmt.Println("13 Verify Signing Shares")
	// fmt.Println("14 U Generation")
	// fmt.Println("15 Signing Phase")
	// fmt.Println("16 Combination Phase")
	// fmt.Println("17 Key Refresh")
	// fmt.Println("2. PreSigning")
	// fmt.Println("3. Signing")
	// fmt.Println("4. Key Refresh")

	// fmt.Scan(&choice)
	time.Sleep(time.Second)
	// fmt.Println(choice)
	// switch choice {
	// case 1:
	time.Sleep(time.Second * 2)
	// Generation of Elgamal Keys
	ESK, EPK := Elgamal_KeyGen()

	fmt.Println(" \n ")
	fmt.Println("Elgamal Public Key:")
	fmt.Println(&EPK)
	fmt.Println("Elgamal Secret Key:")
	fmt.Println(&ESK)
	fmt.Printf("\n")

	//Generating Schnorr Public and  Secret Key
	SSK, SPK := Preprocessing()
	fmt.Println("Schnorr Public Key:")
	fmt.Println(SPK)
	fmt.Println("Schnorr Secret Key:")
	fmt.Println(SSK)
	fmt.Printf("\n")

	//storing the schnorr secret key to Prvate Folder
	os.MkdirAll("Data/"+peer_number, os.ModePerm)
	f1, _ := os.OpenFile("Data/"+peer_number+"/SSK.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	encoding.WriteHexScalar(curve, f1, SSK)
	f1.Close()
	f1, _ = os.OpenFile("Data/"+peer_number+"/SPK.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	encoding.WriteHexPoint(curve, f1, SPK)
	f1.Close()
	//storing the Elgamal Secret key to private folder

	f3, _ := os.OpenFile("Data/"+peer_number+"/ESK.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	fmt.Fprintf(f3, "%x", ESK.Bytes())

	//Storing the Elgamal Public Key to Private folder
	f3, _ = os.OpenFile("Data/"+peer_number+"/EPK.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	fmt.Fprintf(f3, "%x", EPK.ToAffineCompressed())

	// case 2:
	//Reading Elgamal Public key of peer for broadcasting
	file, _ := ioutil.ReadFile("Data/" + peer_number + "/EPK.txt")
	status_struct.Phase = 1
	log.Println(peer_details_list)
	send_data(peer_details_list, string(file), strconv.Itoa(my_index+1), protocolID)
	wait_until(1)

	//VERIFICATION FOR ELGAMAL KEYS AS WELL

	// case 3:
	/****************Commiting SSK and Broadcasting KGC  ****************/
	//Reading SSK
	path := "Data/" + peer_number + "/SSK.txt"
	f2, _ := os.Open(path)
	f_2, _ := encoding.ReadHexScalar(curve, f2)
	f2.Close()

	//commiting SSK
	Commitment(f_2, "hello world", peer_number)
	//Broadcasting KGC
	fmt.Println("Broadcasting KGC values ....")
	fmt.Println("")
	fmt.Println("Broadcasting Signature_S ....")
	f, _ := os.ReadFile("Commitment/" + peer_number + "/KGC/Signature_S" + ".txt")
	status_struct.Phase = 2
	send_data(peer_details_list, string(f), "Signature_S", protocolID)
	wait_until(2)
	fmt.Println("Broadcasting PubKey ....")
	f11, _ := os.ReadFile("Commitment/" + peer_number + "/KGC/PubKey" + ".txt")
	status_struct.Phase = 3
	fmt.Println("-->", string(f11))
	send_data(peer_details_list, string(f11), "PubKey", protocolID)
	wait_until(3)

	fmt.Println("Broadcasting Message ....")
	f31, _ := os.ReadFile("Commitment/" + peer_number + "/KGC/Message" + ".txt")
	status_struct.Phase = 4
	send_data(peer_details_list, string(f31), "Message", protocolID)
	wait_until(4)

	fmt.Println("Broadcasting KGD values ....")

	f4, _ := os.ReadFile("Commitment/" + peer_number + "/KGD" + ".txt")
	status_struct.Phase = 5
	send_data(peer_details_list, string(f4), "KGD", protocolID)
	wait_until(5)

	// case 4:

	var i int64
	//Recieving KGC from peers
	for i = 1; i <= int64(Peer_Count); i++ {
		if i == int64(my_index+1) {
			continue
		}
		Recieve_KGC(strconv.Itoa(int(i)))
	}
	//Recieving KGD from peers
	for i = 1; i <= int64(Peer_Count); i++ {
		if i == int64(my_index+1) {
			continue
		}
		Recieve_KGD(strconv.Itoa(int(i)))
	}

	//Decomiting Values
	for i = 1; i <= int64(Peer_Count); i++ {
		if i == int64(my_index+1) {
			continue
		}
		y_j := Decommitment_j(strconv.Itoa(int(i)))
		//if Decomitment failed
		if y_j == "Invalid" {
			fmt.Printf("Peer %s commited Wrong Values Process Aborting \n", strconv.Itoa(int(i)))
			//break
			//If Decomitment is successful
		} else {
			fmt.Printf("Peer %d Successfully Commited his values \n", i)
			fmt.Printf("Recieved Value from decommitment module is %s \n", y_j)
			fmt.Printf("\n")
		}
	}

	// case 5:

	//Reading Schnorr Secret key of current Peer
	f22, _ := os.Open("Data/" + peer_number + "/SSK.txt")
	SSK, _ = encoding.ReadHexScalar(curve, f22)
	f22.Close()

	//Setting the schnorr secret key, i.e the constant term of Polynomial generated
	// Set_secret(SSK)

	poly := []kyber.Scalar{}  // to store coefficients
	share := []kyber.Scalar{} // to store share
	alphas := []kyber.Point{} // to store alphas

	// var i int64

	for i = 0; i < T; i++ {
		poly = append(poly, curve.Scalar().Zero())
	}

	for i = 0; i < T; i++ {
		alphas = append(alphas, curve.Point().Null())
	}

	for i = 1; i <= int64(Peer_Count); i++ {
		share = append(share, curve.Scalar().Zero())
	}

	// to generate coefficients of the polynomial
	Generate_Polynomial_coefficients(T, poly, peer_number, SSK, "vss/"+peer_number)
	// Generating the shares and storing in share array
	Generate_share(int64(Peer_Count), T, poly, share, peer_number, "vss/"+peer_number)
	//Generating Alphas
	Generate_Alphas(T, alphas, poly, peer_number, "vss/"+peer_number)
	//
	//below commented is for KEY REFRESH
	// Path := "vss/" + peer_number
	// for i = 0; i < T; i++ {
	// 	path1 := Path + "/Initial_coeff" + strconv.Itoa(int(i)) + ".txt"
	// 	file, e2 := os.Create(path1)
	// 	if e2 != nil {
	// 		panic(e2)
	// 	}
	// 	encoding.WriteHexScalar(curve, file, poly[i])
	// 	file.Close()
	// 	//Intializing Curr coeff
	// 	path1 = Path + "/Curr_coeff" + strconv.Itoa(int(i)) + ".txt"
	// 	file, e2 = os.Create(path1)
	// 	if e2 != nil {
	// 		panic(e2)
	// 	}
	// 	encoding.WriteHexScalar(curve, file, poly[i])
	// 	file.Close()
	// }
	// path1 := Path + "/Curr_Share"
	// for i = 1; i <= int64(Peer_Count); i++ {
	// 	file, _ := os.Create(path1 + strconv.Itoa(int(i)) + ".txt")
	// 	encoding.WriteHexScalar(curve, file, share[i-1]) //writing share to the file
	// 	file.Close()
	// }
	//Broadcasting alphas
	status_struct.Phase = 6
	for i = 0; i < T; i++ {

		send_data(peer_details_list, alphas[i].String(), fmt.Sprint(i), protocolID)

	}
	//Waiting for other peers to get to same phase
	wait_until(6)

	//Receiving alphas from other peers
	Recieve_Alphas(int64(Peer_Count), peer_number, T)

	// case 6:
	status_struct.Phase = 7

	fmt.Println("ENCRYPTING SHARES & Broadcasting")

	//Defining Elgamal Curve
	elg_curve := Setup()

	//Reading Sender's elgamal Public key

	path13 := "Data/" + peer_number
	Sender_EPK_file, _ := ioutil.ReadFile(path13 + "/EPK.txt")
	Sender_EPK_file, _ = hex.DecodeString(string(Sender_EPK_file))
	Sender_EPK, _ := elg_curve.Point.FromAffineCompressed(Sender_EPK_file)

	//Reading Sender's elgamal Secret Key
	Sender_ESK_file, _ := ioutil.ReadFile(path13 + "/ESK.txt")
	Sender_ESK_file, _ = hex.DecodeString(string(Sender_ESK_file))
	Sender_ESK, _ := elg_curve.Scalar.SetBytes(Sender_ESK_file)

	//Path to vss generated parameters
	path3 := "vss/" + peer_number

	// var i int64
	for i = 1; i <= int64(Peer_Count); i++ {
		if i == int64(my_index+1) {
			continue
		}

		_f, err := os.Open(path3 + "/Indivisual_Share" + strconv.Itoa(int(i)) + ".txt")
		if err != nil {
			panic(err)
		}

		//share for ith user generated by current peer
		share, _ := encoding.ReadHexScalar(curve, _f)

		//Reading Elgamal Public key of ith user
		data, err := os.ReadFile("Broadcast/" + strconv.Itoa(int(i)) + "/EPK.txt")
		if err != nil {
			fmt.Println(err)
		}
		data, _ = hex.DecodeString(string(data))
		elg_curve := Setup()
		EPK_receiver, _ := elg_curve.Point.FromAffineCompressed(data)

		// fmt.Println("Sender_ESK:", Sender_ESK)
		// fmt.Println("Sender_EPK:", Sender_EPK)
		// fmt.Println("EPK_receiver:", EPK_receiver)
		//encoding.WriteHexScalar(curve, file, shares[i])
		// toEncrypt, _ := encoding.ScalarToStringHex(curve, share)

		//Share to encrypt(in string format)
		toEncrypt := share.String()
		fmt.Println("TO ENCRYPT:", toEncrypt)

		//Ecryption using( current peer's Secret key,current peer's public key, ith users(receivers) public key)
		C1, C2, C3, _ := AuthEncryption(toEncrypt, Sender_ESK, Sender_EPK, EPK_receiver)

		// // CHECKINIG ENC DEC
		// os.MkdirAll("C's/"+peer_number+"/"+fmt.Sprint(i), 0755)
		// f1, _ := os.Create("C's/" + peer_number + "/" + fmt.Sprint(i) + "/C1.txt")
		// f2, _ := os.Create("C's/" + peer_number + "/" + fmt.Sprint(i) + "/C2.txt")
		// f3, _ := os.Create("C's/" + peer_number + "/" + fmt.Sprint(i) + "/C3.txt")

		// _, _ = fmt.Fprintf(f1, "%x", C1.ToAffineCompressed())
		// _, _ = fmt.Fprint(f2, string(C2))
		// _, _ = fmt.Fprint(f3, string(C3))
		// curr_coef-> Updated
		// coeff-> Addend
		// initial-> before any refresh
		// //Trying
		// data, _ = os.ReadFile("Data/" + strconv.Itoa(int(i)) + "/ESK.txt")
		// data, _ = hex.DecodeString(string(data))
		// ESK_receiver, _ := elg_curve.Scalar.SetBytes(data)

		// share_decry, err := AuthDecryption(C1, C2, C3, Sender_EPK, EPK_receiver, ESK_receiver)
		// if err != nil {
		// 	fmt.Println("Error Decrypting in ECNRYP")
		// }

		// original := share.String()
		// fmt.Println("Decrypt share of:", peer_number, "for :", fmt.Sprint(i), share_decry, "\n", "OriginalMessage:", original, "\n")

		// path2 := "CHECKSHARE/" + peer_number + "/Shares/share" + fmt.Sprint(i) + ".txt"
		// os.MkdirAll("CHECKSHARE/"+peer_number+"/Shares/", os.ModePerm)
		// file, _ := os.OpenFile(path2, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
		// _, _ = fmt.Fprint(file, share_decry)
		// fmt.Printf("SENDER EPK: %x\n", Sender_EPK.ToAffineCompressed())
		// fmt.Printf("RECEIVER EPK: %x\n", EPK_receiver.ToAffineCompressed())
		// fmt.Printf("ESK_receiver: %x\n", ESK_receiver.Bytes())

		//
		//// UPTO HERE

		//C := fmt.Sprintf("%x", C1.ToAffineCompressed()) + "||" + string(C2) + "||" + string(C3)

		//Broadcasting the encrypted values.
		var j = 1
		for j = 1; j <= 3; j++ {
			if j == 1 {

				send_data(peer_details_list, fmt.Sprintf("%x", C1.ToAffineCompressed()), fmt.Sprint(i)+","+fmt.Sprint(j), protocolID)
			}
			if j == 2 {
				send_data(peer_details_list, string(C2), fmt.Sprint(i)+","+fmt.Sprint(j), protocolID)

			}
			if j == 3 {
				send_data(peer_details_list, hex.EncodeToString(C3), fmt.Sprint(i)+","+fmt.Sprint(j), protocolID)

			}
		}
		//fmt.Println(C, "\n\n")

		//send_data( peer_details_list, C, fmt.Sprint(i), protocolID)

	}
	wait_until(7)

	fmt.Println("Decrypting Shares")
	//Setting up elgamal curve
	// elg_curve := Setup()

	//Reading elgamal Public key of current peer
	path4 := "Data/" + peer_number
	// Reciever_EPK_file, _ := ioutil.ReadFile(path + "/EPK.txt")
	// Reciever_EPK_file, _ = hex.DecodeString(string(Reciever_EPK_file))
	// Reciever_EPK, _ := elg_curve.Point.FromAffineCompressed(Reciever_EPK_file)

	Reciever_EPK, _ := Get_EPK(path4 + "/EPK.txt")

	//Reading elgamal Secret key of current peer
	// Reciever_ESK_file, _ := ioutil.ReadFile(path + "/ESK.txt")
	// Reciever_ESK_file, _ = hex.DecodeString(string(Reciever_ESK_file))
	// Reciever_ESK, _ := elg_curve.Scalar.SetBytes(Reciever_ESK_file)

	Reciever_ESK, _ := Get_ESK(path4 + "/ESK.txt")

	// var i int64
	for i = 1; i <= int64(Peer_Count); i++ {
		if i == int64(my_index+1) {
			continue
		}

		//Reading Elgamal Public key of Sender(i)
		path2 := "Broadcast/" + fmt.Sprint(i) + "/EPK.txt" //Get the epk of sender
		// data, err := ioutil.ReadFile(path2)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// data, _ = hex.DecodeString(string(data))
		// Sender_EPK, _ := elg_curve.Point.FromAffineCompressed(data)
		Sender_EPK, _ := Get_EPK(path2)

		//Reading the Encrypted data sent by ith user to current user from broadcast folder(channel)
		C1_Data, _ := ioutil.ReadFile("Broadcast/" + fmt.Sprint(i) + "/Shares/To" + peer_number + "/C1.txt")
		C2_Data, _ := ioutil.ReadFile("Broadcast/" + fmt.Sprint(i) + "/Shares/To" + peer_number + "/C2.txt")
		C3_Data, _ := ioutil.ReadFile("Broadcast/" + fmt.Sprint(i) + "/Shares/To" + peer_number + "/C3.txt")

		// C1_Data, _ := ioutil.ReadFile("C's/" + fmt.Sprint(i) + "/" + peer_number + "/C1.txt")
		// C2_Data, _ := ioutil.ReadFile("C's/" + fmt.Sprint(i) + "/" + peer_number + "/C2.txt")
		// C3_Data, _ := ioutil.ReadFile("C's/" + fmt.Sprint(i) + "/" + peer_number + "/C3.txt")

		//Changing the data read into C1,C2,C3 format
		data, _ := hex.DecodeString(string(C1_Data))
		C1, _ := elg_curve.Point.FromAffineCompressed(data)
		C2 := C2_Data
		C3, _ := hex.DecodeString(string(C3_Data))
		// C3 := C3_Data

		//Decryption of shares(C1,C2,C3)
		time.Sleep(time.Second * 5)
		share, err := AuthDecryption(C1, C2, C3, Sender_EPK, Reciever_EPK, Reciever_ESK)
		fmt.Println("DECRPYPteD:", share)

		// fmt.Printf("SENDER EPK: %x\n", Sender_EPK.ToAffineCompressed())
		// fmt.Printf("RECEIVER EPK: %x\n", Reciever_EPK.ToAffineCompressed())
		// fmt.Printf("ESK_receiver: %x\n", Reciever_ESK.Bytes())
		if err != nil {
			fmt.Println("Error Decrypting")
		}

		//Saving the decrypted message into the received folder of current user
		path2 = "Received/" + peer_number + "/Shares/share" + fmt.Sprint(i) + ".txt"
		os.MkdirAll("Received/"+peer_number+"/Shares/", os.ModePerm)
		file, _ := os.OpenFile(path2, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
		_, _ = fmt.Fprint(file, share)

	}

	// case 8:

	fmt.Println(peer_number, "Verifying Shares")

	G := Verify_Share(peer_number, int64(Peer_Count), int64(T), false)
	fmt.Println("Private Key Share:", G.String(), "\n")
	path5 := "Received/" + peer_number
	os.MkdirAll(path5, os.ModePerm)
	file5, _ := os.Create(path5 + "/G.txt")
	encoding.WriteHexScalar(curve, file5, G)

	if verify_GK(int64(Peer_Count), T) {
		fmt.Println("VERIFIED G")
	} else {
		fmt.Println("NOT VERIFIED G")
	}
	//BROADCAST GROUP PUBLIC KEY

	//G-> input to sign t unknwn
	GK := Get_Group_Key(int64(Peer_Count))
	file5, _ = os.Create(path5 + "/GroupKey.txt")
	encoding.WriteHexPoint(curve, file5, GK)
	fmt.Println("GROUP KEY:", GK.String())
	//1 keygen
	//2 singing
	time.Sleep(time.Second * 2)
	fmt.Println("********************KEYGEN COMPLETED*********************")
	Presigning_T_Unknown(peer_number, int64(Peer_Count))

	// case 2:
	// status_struct.Phase = 8

	// fmt.Println("******************************************PRESIGNING PHASE STARTED *******************************************")
	// // var r_i kyber.Scalar
	// var U_i kyber.Point
	// U_i = curve.Point().Null()
	// r_i := curve.Scalar().Pick(curve.RandomStream())
	// U_i = U_i.Mul(r_i, g)
	// //U_i, r_i := Setup_Keys(T, int64(Peer_Count), peer_number, g)
	// os.MkdirAll("Data/"+peer_number+"/Signing/", os.ModePerm)
	// file, _ := os.Create("Data/" + peer_number + "/Signing/r_i.txt")
	// file2, _ := os.Create("Data/" + peer_number + "/Signing/U_i.txt")

	// encoding.WriteHexScalar(curve, file, r_i)
	// encoding.WriteHexPoint(curve, file2, U_i)

	// U_i_sending, _ := os.ReadFile("Data/" + peer_number + "/Signing/U_i.txt")

	// send_data(peer_details_list, string(U_i_sending), "U_i", protocolID)
	// wait_until(8)

	// //	Peer_Commitment(U_i, r_i, peer_number, Peer_Count)

	// fmt.Println("Commiting Signing r_i")
	// file, _ = os.Open("Data/" + peer_number + "/Signing/r_i.txt")
	// file2, _ = os.Open("Data/" + peer_number + "/Signing/U_i.txt")

	// r_i, err := encoding.ReadHexScalar(curve, file)
	// U_i, err2 := encoding.ReadHexPoint(curve, file2)

	// if err != nil || err2 != nil {
	// 	fmt.Println("error occured")
	// }

	// fmt.Println("r_i->", r_i.String(), "U_i->", U_i.String())
	// // path := "Data/" + peer_number + "/SSK.txt"
	// // f2, _ := os.Open(path)
	// // f_2, _ := encoding.ReadHexScalar(curve, f2)
	// // f2.Close()
	// //commiting SSK
	// Commitment_sign(r_i, U_i.String(), peer_number)
	// //Broadcasting KGC

	// // Helper.Broadcast_KGC((peer_number))

	// fmt.Println("Broadcasting KGC values ....")

	// fmt.Println("")
	// fmt.Println("Broadcasting Signature_S ....")
	// f, _ := os.ReadFile("Commitment/Signing/" + peer_number + "/KGC/Signature_S" + ".txt")
	// status_struct.Phase = 9
	// send_data(peer_details_list, string(f), "Signature_S", protocolID)
	// // wait_until(9)

	// fmt.Println("Broadcasting PubKey ....")
	// f1, _ := os.ReadFile("Commitment/Signing/" + peer_number + "/KGC/PubKey" + ".txt")
	// status_struct.Phase = 10
	// fmt.Println("-->", string(f1))

	// send_data(peer_details_list, string(f1), "PubKey", protocolID)
	// // wait_until(10)

	// fmt.Println("Broadcasting Message ....")
	// f3, _ := os.ReadFile("Commitment/Signing/" + peer_number + "/KGC/Message" + ".txt")
	// status_struct.Phase = 11
	// send_data(peer_details_list, string(f3), "Message", protocolID)
	// // wait_until(11)

	// fmt.Println("Broadcasting KGD values ....")

	// f4, _ := os.ReadFile("Commitment/Signing/" + peer_number + "/KGD" + ".txt")
	// status_struct.Phase = 12
	// send_data(peer_details_list, string(f4), "KGD", protocolID)
	// wait_until(12)

	// // case 11:
	// var i int64
	// //Recieving KGC from peers
	// for i = 0; i <= int64(Peer_Count); i++ {
	// 	if i == int64(my_index) {
	// 		continue
	// 	}
	// 	Recieve_KGC_sign(strconv.Itoa(int(i)))
	// }
	// //Recieving KGD from peers
	// for i = 0; i <= int64(Peer_Count); i++ {
	// 	if i == int64(my_index) {
	// 		continue
	// 	}
	// 	Recieve_KGD_sign(strconv.Itoa(int(i)))
	// }

	// //Decomiting Values
	// for i = 0; i <= int64(Peer_Count); i++ {
	// 	if i == int64(my_index) {
	// 		continue
	// 	}
	// 	y_j := Decommitment_j_sign(strconv.Itoa(int(i)))
	// 	if y_j == "Invalid" {
	// 		fmt.Printf("Peer %s commited Wrong Values Process Aborting \n", strconv.Itoa(int(i)))
	// 		//break
	// 	} else {
	// 		fmt.Printf("Peer %d Successfully Commited his values \n", i)
	// 		fmt.Printf("Recieved Value from decommitment module is %s \n", y_j)
	// 		fmt.Printf("\n")
	// 	}
	// }

	// // case 12:
	// //vss k=threshold
	// // f2, _ := os.Open("Received/Signing/" + peer_number + "/G.txt")
	// // x_i, _ := encoding.ReadHexScalar(curve, f2)
	// // f2.Close()
	// // x_i := curve.Scalar().Pick(curve.RandomStream())

	// //Set_secret_sign(x_i)

	// poly := []kyber.Scalar{}  // to store coefficients
	// share := []kyber.Scalar{} // to store share
	// alphas := []kyber.Point{} // to store alphas

	// // pt:=curve.Scalar()
	// // pt.SetBytes()
	// // var i int64

	// for i = 0; i < T; i++ {
	// 	poly = append(poly, curve.Scalar().Zero())
	// }

	// for i = 0; i < T; i++ {
	// 	alphas = append(alphas, curve.Point().Null())
	// }

	// for i = 0; i <= int64(Peer_Count); i++ {
	// 	share = append(share, curve.Scalar().Zero())
	// }

	// // to generate coefficients of the polynomial         //r_i
	// Generate_Polynomial_coefficients(T, poly, peer_number, r_i, "vss/Signing/"+peer_number)
	// // fmt.Println("COFFE", poly[0].String(), "\n", poly[1].String(), "\n")

	// Generate_share(int64(Peer_Count), T, poly, share, peer_number, "vss/Signing/"+peer_number)
	// // fmt.Println("SHARES", share[0].String(), "\n", share[1].String(), "\n")

	// //Generating Alphas
	// Generate_Alphas(T, alphas, poly, peer_number, "vss/Signing/"+peer_number)
	// // fmt.Println("ALPHAS", alphas[0].String(), "\n", alphas[1].String(), "\n")

	// //Broadcasting alphas

	// status_struct.Phase = 13
	// for i = 0; i < T; i++ {
	// 	send_data(peer_details_list, alphas[i].String(), fmt.Sprint(i), protocolID)
	// }

	// wait_until(13)
	// Recieve_Alphas_sign(int64(Peer_Count), peer_number, T)

	// //Broadcasting Share
	// status_struct.Phase = 14
	// paths := "vss/Signing/" + peer_number
	// for i = 0; i <= int64(Peer_Count); i++ {
	// 	_f, _ := os.Open(paths + "/Indivisual_Share" + strconv.Itoa(int(i)) + ".txt")
	// 	shares, _ := encoding.ReadHexScalar(curve, _f)
	// 	tosend := shares.String()
	// 	fmt.Println("TO SeND:", tosend)
	// 	send_data(peer_details_list, tosend, fmt.Sprint(i), protocolID)
	// }
	// wait_until(14)

	// //Receiving Sign shares
	// Recieve_Share_sign(peer_number, int64(Peer_Count))
	// // var p int
	// // for p = 0; p <= Peer_Count; p++ {
	// // 	if strconv.Itoa(p) == peer_number {
	// // 		continue
	// // 	}
	// // 	alphas2 := []kyber.Point{} // to store alphas
	// // 	for i = 0; i < T; i++ {
	// // 		alphas2 = append(alphas2, curve.Point().Null())
	// // 	}

	// // 	var q int
	// // 	for q = 0; q < int(T); q++ {
	// // 		fx, err := os.Open("vss/Signing/" + strconv.Itoa(int(p)) + "/alpha" + strconv.Itoa(int(q)) + ".txt")
	// // 		if err != nil {
	// // 			panic(err)
	// // 		}
	// // 		fx_p, _ := encoding.ReadHexPoint(curve, fx)
	// // 		alphas2[q] = fx_p
	// // 	}
	// // 	fx2, errx := os.Open("vss/Signing/" + strconv.Itoa(int(p)) + "/Indivisual_Share" + peer_number + ".txt")
	// // 	if errx != nil {
	// // 		fmt.Print("ERROR ERRX")
	// // 	}
	// // 	my_share, _ := encoding.ReadHexScalar(curve, fx2)
	// // 	check(peer_number, my_share, T, alphas2)

	// // }

	// // case 13:
	// fmt.Println("Verifying Signing Shares")

	// path := "Received/Signing/" + peer_number + "/R_i.txt"
	// R_i := Verify_Share(peer_number, int64(Peer_Count), T, true)
	// file3, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	// if err != nil {
	// 	panic(err)
	// }
	// encoding.WriteHexScalar(curve, file3, R_i)
	// file3.Close()

	// // case 14:
	// var U kyber.Point
	// U = curve.Point().Null()
	// fmt.Println("U_i's : ")
	// // var i int
	// for i = 0; i <= int64(Peer_Count); i++ {
	// 	if i == int64(my_index) {
	// 		path := "Data/" + fmt.Sprint(i) + "/Signing/U_i.txt"
	// 		file, err := os.Open(path)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		U_i, e1 := encoding.ReadHexPoint(curve, file)
	// 		if e1 != nil {
	// 			panic(e1)
	// 		}
	// 		fmt.Println(U_i)
	// 		if U.Equal(curve.Point().Null()) {
	// 			U = U_i
	// 		} else {
	// 			U.Add(U, U_i)
	// 		}
	// 	} else {
	// 		peer := strconv.Itoa(int(i))
	// 		U_i := Read_Ui(peer)
	// 		fmt.Println(U_i)

	// 		if U.Equal(curve.Point().Null()) {
	// 			U = U_i
	// 		} else {
	// 			U.Add(U, U_i)
	// 		}
	// 	}
	// }
	// // //U,U_i,R_i is generated for each peer above
	// fmt.Println("U:")
	// fmt.Println(U)

	// os.MkdirAll("Data/"+peer_number+"/Signing/", os.ModePerm)
	// file, _ = os.Create("Data/" + peer_number + "/Signing/U.txt")
	// encoding.WriteHexPoint(curve, file, U)
	// choice = 3
	// time.Sleep(time.Second * 2)
	// case 3:
	// fmt.Printf("********************************************* SIGNING PHASES STARTED ******************************************\n")

	// file, _ := os.Open("Data/" + peer_number + "/Signing/r_i.txt")
	// R_i, _ := encoding.ReadHexScalar(curve, file)

	// file, _ = os.Open("Received/Signing/" + peer_number + "/G.txt")
	// x_i, _ := encoding.ReadHexScalar(curve, file)

	// file, _ = os.Open("Data/" + peer_number + "/Signing/U.txt")
	// U, _ := encoding.ReadHexPoint(curve, file)

	// Message := "Hello Aasif Sign"
	// V_i := Signing_T_Unkown(U, x_i, R_i, Message)

	// file, _ = os.Create("Data/" + peer_number + "/Signing/V_i.txt")
	// encoding.WriteHexScalar(curve, file, V_i)

	// //Broadcasting V_i
	// status_struct.Phase = 15

	// tosend, _ := encoding.ScalarToStringHex(curve, V_i)
	// send_data(peer_details_list, tosend, "V_i", protocolID)

	// wait_until(15)

	// status_struct.Phase = 16

	// tosend, _ = encoding.PointToStringHex(curve, U)
	// send_data(peer_details_list, tosend, "U", protocolID)

	// wait_until(16)
	// choice = 4
	// time.Sleep(time.Second * 2)
	// case 4:
	// 	// fmt.Println("************ COMBINATION PHASE ****************")
	// 	// T_arr := [...]int{1}
	// 	// combine_T_Unknown(T_arr[:], peer_number)
	// 	// choice = 291
	// case 5: //verify single
	// case 17:
	// 	Key_Refresh(T, int64(Peer_Count), peer_number, peer_details_list, protocolID)
	// case 29:
	// 	os.Exit(0)
	// case 291:
	// 	break
	// }

}

// //Adding random compute time to show mismatch send times
// a := rand.Intn(10)
// log.Println("Seconds:", a)
// time.Sleep(time.Duration(a) * time.Second)

// //Send kgc_i values
// send_data(peer_details_list, tosend, "U", protocolID)

// status_struct.Phase = 2
// send_data(peer_details_list, "test kgc_i value", "kgc_j", protocolID)
// wait_until(2)

// //Send kgd_i values
// status_struct.Phase = 3
// send_data(peer_details_list, "test kgd_i value", "kgd_j", protocolID)
// wait_until(3)

//?!?!VSS??!?!

// func SEND_DATA(p2p *peer_id peer.ID, value string, name string, protocolID protocol.ID) {

// 	peer_num, _ := strconv.Atoi(string(peer_id))
// 	peer_num = peer_num + len(p2p.Peers)/2
// 	message_send := message{
// 		Phase: status_struct.Phase,
// 		Name:  name,
// 		Value: value,
// 		To:    peer_map[string(peer_id)],
// 	}

// 	//log.Println(peer_map[item])
// 	//log.Println("Sending to index")
// 	s, err := p2p.Host.NewStream(p2p.Ctx, peer_map[peer_id.String()], protocolID)
// 	if err != nil {
// 		log.Println(peer_map[string(peer_id)])
// 		log.Println(err, "Connecting to send message error")
// 		return
// 	}

// 	b_message, err := json.Marshal(message_send)
// 	if err != nil {
// 		log.Println(err, "Error in jsonifying data")
// 		return
// 	}
// 	_, err = s.Write(append(b_message, '\n'))

// 	if err != nil {
// 		log.Println(err, "Sending message erorr")
// 		return
// 	}

// }
func send_data(peer_details_list []string, value string, name string, protocolID protocol.ID) {

	log.Println("Sending phase:", status_struct.Phase)
	for i, item := range peer_details_list {

		// log.Println(peer_details_list)
		// log.Println(i, p2p.Host_ip, item, my_index)
		//log.Println(i, my_index, p2p.Host_ip, item)
		if i == my_index {
			continue
		}
		addr, _ := multiaddr.NewMultiaddr(item)
		peer_info, err := peer.AddrInfoFromP2pAddr(addr)

		if err != nil {
			panic(err)
		}

		// peer_num, _ := strconv.Atoi(item)
		// peer_num = peer_num + len(p2p.Peers)/2
		message_send := message{
			Phase: status_struct.Phase,
			Name:  name,
			Value: value,
			To:    peer_map[item],
		}

		s, err := p2p.Host.NewStream(p2p.Ctx, peer_info.ID, protocolID)
		if err != nil {
			log.Println(peer_map[item])
			log.Println(err, "Connecting to send message error")
			return
		}

		b_message, err := json.Marshal(message_send)
		if err != nil {
			log.Println(err, "Error in jsonifying data")
			return
		}
		_, err = s.Write(append(b_message, '\n'))

		if err != nil {
			log.Println(err, "Sending message erorr")
			return
		}

	}

}

func wait_until(phase int) {
	for {
		flag := 0
		for i, item := range peer_details_list {
			item = strings.Split(item, "/")[len(strings.Split(item, "/"))-1]
			if i == my_index {
				continue
			}
			if phase != receive_peer_phase[item] {
				// if phase > receive_peer_phase[item] {
				// 	// Resend value to 'item'
				// }
				flag = 1
				// log.Println("heres why: ", receive_peer_phase[item])
			}
			if phase != sent_peer_phase[item] {
				flag = 1
				// log.Println("heres why: ", sent_peer_phase[item])
			}

		}
		if flag == 1 {

			time.Sleep(time.Millisecond * 10)
			// log.Println(flag, phase, receive_peer_phase, sent_peer_phase)
			flag = 0
			continue
		}
		fmt.Println("Returning from phase ", phase)
		// time.Sleep(time.Second)
		// log.Println(phase, receive_peer_phase, sent_peer_phase)
		return

	}
}
func Wait_until_for_sign(phase int, t int) {
	for {
		flag := 0
		for i, item := range peer_details_list {
			item = strings.Split(item, "/")[len(strings.Split(item, "/"))-1]
			if i == my_index {
				continue
			}
			if phase != receive_peer_phase[item] {
				// if phase > receive_peer_phase[item] {
				// 	// Resend value to 'item'
				// }
				flag += 1
				// log.Println("heres why: ", receive_peer_phase[item])
			} else if phase != sent_peer_phase[item] {
				flag += 1
				// log.Println("heres why: ", sent_peer_phase[item])
			}

		}
		if flag != len(peer_details_list)-t {

			time.Sleep(time.Microsecond * 5)
			// log.Println(flag, phase, receive_peer_phase, sent_peer_phase)
			flag = 0
			continue
		}
		fmt.Println("Returning from phase ", phase)
		// time.Sleep(time.Second)
		// log.Println(phase, receive_peer_phase, sent_peer_phase)
		return

	}
}
func GeneratePrime(size int) *big.Int {
	prime, err := rand.Prime(rand.Reader, size)
	if err != nil {
		fmt.Println(err)
	}
	return prime
}
