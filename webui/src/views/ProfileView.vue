<script>

const READER = new FileReader();
const STATOBAN = ["BAN", "UNBAN"];
const STATOFOLLOW = ["FOLLOW", "UNFOLLOW"];

export default{
  data(){
    return{
      newUsername : "",

      errormsg: "",
      
      visitor : {
        visitorID : Number(localStorage.getItem('identifier')),
        visitorNome : localStorage.getItem('username'),
        visitorSeguaci : [],
        visitorSeguiti : [],
        visitorBannati : [],
        hasBanned : null,
        isFollowing : null,
      },

      profilo : {
        ID : Number(localStorage.getItem('IDCercato')),
        nome : "",
        seguaci : [],
        seguiti : [],
        bannati : [],
        fotoPubblicate : []
      }
    }
  },
  methods: {
    async refresh(){
      await this.recuperaInfo();
      if (this.visitor.visitorID != this.profilo.ID) { await this.visitorInfo(); }
    },
    ///////////////////////////// FUNZIONI VISITOR ////////////////////////////////////////////
    async visitorInfo(){
      try
      {
        let response = await this.$axios.get('/userProfile/' + this.visitor.visitorID, {headers: {Authorization: "Bearer " + this.visitor.visitorID}});
        
        if (response.data.Followers != null) {this.visitor.visitorSeguaci = response.data.Followers; }
        
        if (response.data.Followings != null) {this.visitor.visitorSeguiti = response.data.Followings; }
        
        if (response.data.Banneds != null) {this.visitor.visitorBannati = response.data.Banneds;}

        this.visitor.isFollowing = this.visitor.visitorSeguiti.includes(this.profilo.ID);

        this.visitor.hasBanned = this.visitor.visitorBannati.includes(this.profilo.ID);
      }
      catch(e)
      {
        this.errormsg = e.toString();
        alert(this.errormsg);
      }
    },
    async gestioneFollow(){
      try
      {
        if (this.visitor.isFollowing) {
          await this.$axios.delete('/userProfile/' + this.visitor.visitorID + '/following/' + this.profilo.ID, {headers: {Authorization: "Bearer " + this.visitor.visitorID}} ); 
          alert("Non segui più " + this.profilo.nome);
        }
        else {
          await this.$axios.post('/userProfile/' + this.visitor.visitorID + '/following/' + this.profilo.ID, {}, {headers: {Authorization: "Bearer " + this.visitor.visitorID}} );
          alert("Ora segui " + this.profilo.nome);
        }
        this.refresh();
      }
      catch(e)
      {
        this.errormsg = e.toString();
        alert(this.errormsg);
      }
    },
    async gestioneBan(){
      try
      {
        if (this.visitor.hasBanned) {
          await this.$axios.delete('/userProfile/' + this.visitor.visitorID + '/banList/' + this.profilo.ID, {headers: {Authorization: "Bearer " + this.visitor.visitorID}} ); 
          alert(this.profilo.nome + " non è più bannato!");
        }
        else {
          await this.$axios.post('/userProfile/' + this.visitor.visitorID + '/banList/' + this.profilo.ID, {}, {headers: {Authorization: "Bearer " + this.visitor.visitorID}} );
          alert("Hai bannato " + this.profilo.nome);
        }
        this.refresh();
      }
      catch(e)
      {
        this.errormsg = e.toString();
        alert(this.errormsg);
      }
    },
    ////////////////////////////// FUNZIONI OWNER //////////////////////////////////////////////
    async recuperaInfo(){
      try
      {
        let response = await this.$axios.get('/userProfile/' + this.profilo.ID, {headers: {Authorization: "Bearer " + this.visitor.visitorID}});
        
        this.profilo.nome = response.data.Username;
        
        if (response.data.Followers != null) {this.profilo.seguaci = response.data.Followers; }
        
        if (response.data.Followings != null) {this.profilo.seguiti = response.data.Followings; }
        
        if (response.data.Banneds != null) {this.profilo.bannati = response.data.Banneds;}
        
        if (response.data.PublishedPhotos != null) {this.profilo.fotoPubblicate = response.data.PublishedPhotos;}
      }
      catch(e)
      {
        this.errormsg = e.toString();
        alert(this.errormsg);
      }
    },
    async uploadPhoto(){
      try
      {
        let foto = document.getElementById('upload');
        READER.readAsArrayBuffer(foto.files[0]);
        READER.onload = async () => {
          let response = await this.$axios.put('/userProfile/${this.visitorID}/publishedPhotos/', READER.result, { headers: {Authorization: "Bearer " + this.visitor.visitorID} });
          alert("Foto pubblicata!");
        };
        this.refresh();
      }
      catch(e)
      {
        this.errormsg = e.toString();
        alert(this.errormsg);
      }
    },
    async mostraDialog(){ document.getElementById("cambiaNome").style.display = "inline"; },
    
    async chiudiDialog(){
      this.newUsername = "";
      document.getElementById("cambiaNome").style.display = "none"; 
    },
    
    async cambiaUsername(){
      try
      {
        let response = await this.$axios.post('/user/' + this.profilo.ID + '/username', {username: this.newUsername}, { headers: {Authorization: "Bearer " + this.visitor.visitorID} })
        this.profilo.nome = this.newUsername;
        localStorage.setItem('username', this.newUsername);
        this.newUsername = "";
        alert("Username cambiato!");
        this.chiudiDialog();
      }
      catch(e)
      {
        this.errormsg = e.toString();
        alert(this.errormsg);
      }
    },
    async controllaSeguaci(){
      this.$router.push({path: '/userProfile/' + this.profilo.ID + '/following'});
    },
    async controllaSeguiti(){
      this.$router.push({path: '/userProfile/' + this.profilo.ID + '/followeds'});
    },
    async controllaBannati(){
      this.$router.push({path: '/userProfile/'+ this.profilo.ID + '/banList'});
    },
    async tornaHomePage(){
      localStorage.removeItem('IDCercato');
      this.$router.replace('/session');
    }
  },
    computed: {
      controllaBan() {return this.visitor.hasBanned ? STATOBAN[1] : STATOBAN[0];},
      controllaFollow() { return this.visitor.isFollowing ? STATOFOLLOW[1] : STATOFOLLOW[0];}
  },
    mounted(){
      this.refresh();
  }
}
</script>

<template>
  <body>
    <header class=intestazione>
      <h4 @click = "tornaHomePage" style= "padding-left : 10px; font-weight: bold;">Torna alla HomePage</h4>
      <h1 style = "font-weight: bold;">Profilo di {{this.profilo.nome}}</h1>
      <div class=statistiche style= "padding-right : 10px; ">
        <h4 @click = "controllaSeguaci">Seguaci: {{this.profilo.seguaci.length}}</h4>
        <h4 @click = "controllaSeguiti">Seguiti: {{this.profilo.seguiti.length}}</h4>
      </div>
    </header>

    <dialog id = "cambiaNome" close>
      <h3 style = "text-align: center;"> Scrivi il tuo nuovo username </h3>
      <div style = "align-self: center;">
        <input placeholder="Nuovo username" v-model = this.newUsername style = "text-align: center;">
        <button @click = "cambiaUsername" :disabled= "this.newUsername.length > 16 || this.newUsername.length < 3">Cambia</button>
        <button @click = "chiudiDialog" > Chiudi </button>
      </div>
    </dialog>

    <div class = fotoCondivise v-for = "immagine in this.profilo.fotoPubblicate">
      <div class = colonna>
        <img :src = "immagine.File">
      </div>
    </div>

    <div class = footer>
      <div class = ownerActions v-if = "this.profilo.ID == this.visitor.visitorID">
        <h4 style= "padding-left : 100px; font-weight: bold;" @click = "controllaBannati"> Bannati: {{this.profilo.bannati.length}}</h4>
        <input id = "upload" type="file" accept="image/*" @change="uploadPhoto" style= "padding-left : 100px;">
        <h4 style= "padding-right : 100px; font-weight: bold;" @click = "mostraDialog()"> Cambia username </h4>
      </div>

      <div class = visitorActions v-else>
        <button id = banButton style = "font-size: 30px; color:red;" @click = "gestioneBan"> {{controllaBan}} </button>
        <button id = followButton style = "font-size: 30px;" @click = "gestioneFollow"> {{controllaFollow}} </button>
      </div>
    </div>
  </body>
</template>

<style>

  dialog[open]{
    top:  50%;
    left: 50%;
    display: inline;
    transform: translate(-100%, -60%);
  }

  body{
    margin: 0;
    margin-right: 0;
    width:100%;
    background-color: whitesmoke;
  }

  .intestazione{
    width:100%;
    display: flex;
    justify-content: space-between;
    align-items: center;
    cursor: pointer;
    height:100px;

    background-color: rgb(0, 117, 94);

    border-bottom-left-radius: 30px;
    border-bottom-right-radius: 30px;
    border-bottom: 3px solid black;
  }

  .riga{
    padding-top: 300px;
    display: flex;
    flex-wrap: wrap;
  }

  .colonna{
    flex: 1 0 30%;
    padding: 5px;
  }

  .foto{
    width: 100%;
    height:100%;
    border-radius: 10px;
  }

  .footer{
    position:fixed;
    width:100%;

    cursor: pointer;
    height:100px;
    bottom: 0;

    background-color:rgb(0, 117, 94);

    border-top-left-radius: 30px;
    border-top-right-radius: 30px;
    border-top: 3px solid black;
  }

  .ownerActions{
    padding-top: 25px;
    display: flex;
    align-items: center;
    justify-content:space-between;
  }

  .visitorActions{
    padding-top: 25px;
    display: flex;
    align-items: center;
    justify-content:space-evenly;
  }
</style>