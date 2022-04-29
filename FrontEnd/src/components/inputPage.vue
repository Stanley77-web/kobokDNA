<template>
    <div>
    <img src="../assets/kobo2_resized.png" style="position:fixed;bottom:0;right:0">
    <div v-if="mainMenu" class = "box-wrapper">
        <h1 style="margin-bottom: 100px;">
            Main Menu
        </h1>
        <div class = "menu-option">
            <button @click="{ mainMenu = false; testDNA = true}" class = "menu-option-button">Test DNA</button>
            <button @click="{ mainMenu = false; addPenyakit = true}" class = "menu-option-button">Penambahan Penyakit</button>
            <button @click="{ mainMenu = false; searchDatabase = true}" class = "menu-option-button">Search Database</button>
        </div>
    </div>
    
    <div v-if="testDNA" @submit.prevent="tesDNA" action="" class = "box-wrapper">
        <button @click="{mainMenu = true; testDNA = false; reset()}" class="back-button">back</button>
        <h1 style="margin-bottom: 20px;">
            Test DNA
        </h1>
        <div class="input-option-wrapper">
            <div class = "input-option">
                <h3>Nama Pengguna</h3>
                <input type="text" v-model="namaPengguna" class = "menu-option-input" placeholder="<pengguna>">
            </div>
            <div class = "input-option">
                <h3>Upload Sequence</h3>
                <label class = "custom-file-upload">
                    <input type="file" @change="onFileUploaded($event.target.files)" accept=".txt">
                    <a v-if="uploaded">
                        {{namaFile}}
                    </a>
                    <a v-else>
                    Upload Sequence...
                    </a>
                </label>
            </div>
            <div class = "input-option">
                <h3>Prediksi Penyakit</h3>
                <input type="text" v-model="namaPenyakit" class = "menu-option-input" placeholder="<penyakit>">
            </div>
            <div v-if="showResult">
                {{resultTanggal}} - {{resultNamaPengguna}} - {{resultNamaPenyakit}} - {{persentase}}% - {{diagnosis ? "True" : "False"}}
            </div>
            <div v-if="showError" style="color:red">
                {{errorMessage}}
            </div>
        </div>
        <div class = "radio-button">
            <input type="radio" v-model="isKMP" value = "1">KMP
            <input type="radio" v-model="isKMP" value = "0">BM
        </div>
        <div class="submit-div">
            <button @click="tesDNA();showError = false;showResult = false">Submit</button>
        </div>
    </div>

    <div v-if="addPenyakit" @submit.prevent="addPenyakit" action="" class = "box-wrapper">
        <button @click="{mainMenu = true; addPenyakit = false; reset()}" class="back-button">back</button>
        <h1 style="margin-bottom: 20px;">
            Tambah Penyakit
        </h1>
        <div class="input-option-wrapper">
            <div class = "input-option">
                <h3>Nama Penyakit</h3>
                <input type="text" v-model="newPenyakit" class = "menu-option-input" placeholder="<penyakit>">
            </div>
            <div class = "input-option">
                <h3>Upload Sequence</h3>
                <label class = "custom-file-upload">
                    <input type="file" @change="onFileUploaded($event.target.files)" accept=".txt">
                    <a v-if="uploaded">
                        {{namaFile}}
                    </a>
                    <a v-else>
                        Upload Sequence...
                    </a>
                </label>
            </div>
            <div v-if="showResult" style="color:green">
                {{succMessage}}
            </div>
            <div v-if="showError" style="color:red">
                {{errorMessage}}
            </div>
        </div>
        
        <div class="submit-div">
            <button @click="addDisease();showError = false;showResult = false">Submit</button>
        </div>
    </div>


    <div v-if="searchDatabase" action="" class = "box-wrapper">
       <button @click="{mainMenu = true; searchDatabase = false; reset()}" class="back-button">back</button>
        <h1 style="margin-bottom: 20px;">
            Search Predictions
        </h1>
        <div class="input-option-wrapper">
            <div class = "input-option">
                <input type="text" v-model="searchQueryInput" class="database-search-input" @keyup.enter="searchQuery(); showError = false; showResult = false; namaPenyakit = ''; tanggal = ''">
            </div>
            <div v-if="showError" style="color:red">
                {{errorMessage}}
            </div>
        </div>
        <div v-if="queryEntered">
            <div v-if="showResult">
                <ul>
                    <li v-for="item in arrPenyakit.slice(lowIndex,highIndex)" :key="item.index" class="hasil-query">
                        <div class = "data-hasil-query">
                            {{item.date}} - {{item.user_name}} - {{item.disease_name}} - {{item.similarity_score}}% - {{item.status ? "True" : "False"}}
                        </div>
                    </li>
                </ul>
                <div class="next-prev-wrapper">
                    <div v-if="displayPrev">
                        <button class="prev-next-button" @click="prevQuery()">
                            &lt;
                        </button>
                    </div>
                    <div v-if="displayNext">
                        <button class="prev-next-button" @click="nextQuery()">
                            &gt;
                        </button>
                    </div>
                </div>
            </div>
        </div>
    </div>
    </div>
</template>

<script>
// import axios from 'axios';
import axios from 'axios'

export default {
    
    data(){
        return{
            mainMenu: true,
            addPenyakit: false,
            testDNA: false,
            searchDatabase: false,
            uploaded: false,
            queryEntered: false,
            showResult: false,
            showError: false,
            displayPrev: false,
            displayNext: false,
            diagnosis: "False",
            searchQueryInput: "",
            namaPengguna: "",
            namaFile: "",
            namaPenyakit: "",
            newPenyakit: "",
            respPenyakit: "",
            isiFile: "",
            tanggal: "",
            errorMessage: "",
            succMessage: "",
            resultTanggal: "",
            resultNamaPengguna: "",
            resultNamaPenyakit: "",
            persentase: 0,
            isKMP: "1",
            port: "http://localhost:8080",
            arrPenyakit: [{index: 1,namaPenyakit: "a",namaFile: "a"},
                        {index: 2,namaPenyakit: "b",namaFile: "b"},
                        {index: 3,namaPenyakit: "c",namaFile: "c"},
                        {index: 4,namaPenyakit: "d",namaFile: "d"},
                        {index: 5,namaPenyakit: "e",namaFile: "e"},
                        {index: 6,namaPenyakit: "f",namaFile: "f"},
                        {index: 7,namaPenyakit: "g",namaFile: "g"},
                        {index: 8,namaPenyakit: "h",namaFile: "h"},
                        {index: 9,namaPenyakit: "i",namaFile: "i"},
                        {index: 10,namaPenyakit: "j",namaFile: "j"},
                        {index: 11,namaPenyakit: "k",namaFile: "k"},
                        {index: 12,namaPenyakit: "l",namaFile: "l"}
                        ],
            lowIndex: 0,
            highIndex: 10,
            queryLength: 0,
            dataPerPage: 10,
        }
    },
    methods:{
        onFileUploaded(listFile){
            this.uploaded = true;
            this.namaFile = listFile[0].name;
            const fileReader = new FileReader();
            fileReader.onload = (res) =>{
                this.isiFile = res.target.result;
            };
            fileReader.readAsText(listFile[0]);
        },

        reset(){                   //Digunakan untuk mereset semua state dan server json
            this.namaPengguna = "";
            this.namaFile = "";
            this.namaPenyakit = "";
            this.searchQueryInput = "";
            this.uploaded = false;
            this.queryEntered = false;
            this.showResult = false;
            this.showError = false;
            this.diagnosis = "False";
            this.persentase = 0;
            this.newPenyakit = "";
            this.isiFile = "";
            this.tanggal = "";
            this.isKMP = "1";
        },

        tesDNA() {
            var data_pass = {"user_name": this.namaPengguna, "user_DNA_sequence": this.isiFile, "disease_name": this.namaPenyakit, "method": this.isKMP};
            axios({ method: "POST", url: this.port+"/v1/testDNA/test", data: data_pass, headers: {"content-type": "text/plain" } }).then(result => { 

                if (result.data["message"] == "success adding a result") {
                    var query = "?id=" + result.data["id"];
                    axios({ method: "GET", url: this.port+"/v1/testDNA/result"+query, headers: {"content-type": "text/plain" } }).then(result => { 
                    this.resultTanggal = result.data["date"];
                    this.resultNamaPengguna = result.data["user_name"];
                    this.resultNamaPenyakit = result.data["disease_name"];
                    this.persentase = result.data["similarity_score"];
                    this.diagnosis = result.data["status"];
                    this.showResult = true;
                    }).catch( error => {   
                        this.errorMessage = error.response.data["message"];                 
                        this.showError = true;
                })}           

                }).catch( error => {   
                    this.errorMessage = error.response.data["message"];                 
                    this.showError = true;
        })},

        addDisease(){
            var data_pass = {"disease_name": this.newPenyakit, "disease_sequence_DNA": this.isiFile};
            this.addPenyakit = true;

            axios({ method: "POST", url: this.port+"/v1/disease/add", data: data_pass, headers: {"content-type": "text/plain" } }).then(result => { 
                    this.succMessage = result.data["message"];

                    this.showResult = true;
                    }).catch( error => {    
                        this.errorMessage = error.response.data["message"];                 
                        this.showError = true;
            })},

        searchQuery(){
            this.lowIndex = 0;
            this.highIndex = 10;
            if (this.searchQueryInput == "") {
                this.queryEntered = false;
            } else {

                var query = "?query=" + this.searchQueryInput;
                    axios({ method: "GET", url: this.port+"/v1/searching/predictionResult"+query, headers: {"content-type": "text/plain" } }).then(result => { 
                    this.queryEntered = true;
                    this.arrPenyakit = result.data["predictions"];
                    this.queryLength = result.data["predictions"].length;

                    if (this.queryLength > this.dataPerPage) {
                        this.displayNext = true;
                    }

                    this.showResult = true;
                    }).catch( error => {
                        this.errorMessage = error.response.data["message"];                 
                        this.showError = true;      
                })}
                this.displayPrevNext();
            
        },
        nextQuery(){
            if(this.highIndex + this.dataPerPage >this.queryLength && this.queryLength%this.dataPerPage !== 0){
                this.lowIndex = this.queryLength-this.queryLength%this.dataPerPage;
                this.highIndex = this.queryLength;
                
            }else if(this.highIndex + this.dataPerPage <= this.queryLength){
                this.lowIndex = this.lowIndex + this.dataPerPage;
                this.highIndex = this.highIndex + this.dataPerPage;
            }
            this.displayPrevNext()
        },
        prevQuery(){
            if(this.lowIndex - this.dataPerPage < 0){
                this.lowIndex = 0;
                this.highIndex = this.dataPerPage;
            }else if(this.highIndex === this.queryLength && this.queryLength%this.dataPerPage !== 0){
                var x;
                x = this.queryLength % this.dataPerPage;
                this.lowIndex = this.lowIndex - this.dataPerPage;
                this.highIndex = this.highIndex - x;
            }else{
                this.lowIndex = this.lowIndex - this.dataPerPage;
                this.highIndex = this.highIndex - this.dataPerPage;
            }
            this.displayPrevNext()
        },

        displayPrevNext() {
            if (this.lowIndex + this.dataPerPage >= this.queryLength) {
                this.displayNext = false;
            } else {
                this.displayNext = true;
            }
            if (this.lowIndex < this.dataPerPage) {
                this.displayPrev = false;
            } else {
                this.displayPrev = true;
            }
        },
        
        
    }, 

}
</script>

<style>
    h1 {
        text-align: center;
        font-size: 50px;
        font-family: 'Roboto', sans-serif;
        color: rgb(121, 195, 255);
        margin-top: 10px;
        

    }
    body {
        background-image: url('../assets/cloud3.png');
    }
    .box-wrapper{
        width: 60%;
        margin: 20px auto; 
        min-height: 400px;
        min-width: 800px;
        background: white;
        text-align: justify;
        padding: 40px;
        border-radius: 20px;
    }
    .menu-option{
        text-align: center;
    }
    .menu-option button{
        border: 3px solid black;
        background-color: white;
        color: rgb(129, 162, 255);
        padding: 14px 28px;
        font-size: 16px;
        cursor: pointer;
        border-radius: 8px;
        margin: 0px 10px;
        min-width: 240px;
        min-height: 75px;
        border-color: #76c1ff;
        font-family: 'Times New Roman', Times, serif;
        font-size: 25px;
        -webkit-text-stroke-width: 0.2px;
        -webkit-text-stroke-color: black;
    }
    .menu-option-button:hover{
        background-color: #76c1ff;
        color: white;
        -webkit-text-stroke-width: 0px;
    }

    .input-option-wrapper{
        text-align: center;
    }

    .next-prev-wrapper{
        display: flex;
        justify-content: center;
    }

    .input-option{
        display: inline-block;
        margin: 20px;
    }

    .input-option input{
        display: block;
    }
    .input-option h3{
        /* margin-left: 5px; */
        text-align: left;
    }

    .back-button{
        position:absolute;
        border: 3px solid #76c1ff;
        background-color: white;
        color: rgb(129, 162, 255);
        padding: 14px 28px;
        font-size: 16px;
        cursor: pointer;
        border-radius: 8px;
        margin: 0px 10px;
        font-family: 'Times New Roman', Times, serif;
        font-size: 25px;
        -webkit-text-stroke-width: 0.2px;
        -webkit-text-stroke-color: black;
    }
    .back-button:hover{
        background-color: #76c1ff;
        color: white;
        -webkit-text-stroke-width: 0px;
    }

    input[type="file"] {
        display: none;
    }
    input[type="text"] {
        height: 25px;
    }

    .custom-file-upload {
        border: 1px solid #ccc;
        display: inline-block;
        padding: 6px 48px;
        cursor: pointer;
    }

    .database-search-input {
        height: 50px;
        width: 300px;
        max-width: 300px;
        max-height: 50px;
        padding: 0px 15px;
        font-size: 20px;
        text-align:center;
    }
    
    .prev-next-button{
        background-color: #76c1ff;
        color: white;
        border-radius: 10px;
        font-size: 25px;
        margin-left: 5px;
        border: 1px solid black;
    }

    .hasil-query{
        list-style: none;
        padding: 5px;
        border: 2px solid black;
        border-radius: 7px;
        margin: 10px;
    }

    .data-hasil-query{
        font-size: 20px;
        text-align: center;
    }

    .submit-div{
        text-align:center;
    }

    .submit-div button{
        border: 3px solid #76c1ff;
        background-color: white;
        color: rgb(129, 162, 255);
        padding: 14px 28px;
        font-size: 16px;
        cursor: pointer;
        border-radius: 8px;
        margin: 20px 10px;
        font-family: 'Times New Roman', Times, serif;
        font-size: 25px;
        -webkit-text-stroke-width: 0.2px;
        -webkit-text-stroke-color: black;
    }

    .submit-div button:hover{
        background-color: #76c1ff;
        color: white;
        -webkit-text-stroke-width: 0px;
    }

    .radio-button{
        padding-right: 10px;
        text-align: center;
    }
</style>
