<script>

const READER = new FileReader();

export default{
  data(){
    return{
      visitorID : localStorage.getItem('identifier'),
      newUsername : "",
      errormsg: "",

      profilo : {
        ID : localStorage.getItem('IDCercato'),
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
    },
    async recuperaInfo(){
      try
      {
        let response = await this.$axios.get('/userProfile/' + localStorage.getItem('IDCercato'), {headers: {Authorization: "Bearer " + this.visitorID}});
        this.profilo.nome = response.data.Username;
        if (response.data.Followers != null) {this.profilo.seguaci = response.data.Followers; }
        if (response.data.Followings != null) {this.profilo.seguiti = response.data.Followings; }
        if (response.data.Banneds != null) {this.profilo.bannati = response.data.Banneds;}
        if (response.data.PublishedPhotos != null) { this.profilo.fotoPubblicate = response.data.PublishedPhotos; }
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
          let response = await this.$axios.put('/userProfile/${this.visitorID}/publishedPhotos/', READER.result, { headers: {Authorization: "Bearer " + this.visitorID} });
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
    async mostraDialog(){ document.getElementById("cambiaNome").show(); },
    
    async chiudiDialog(){ document.getElementById("cambiaNome").close(); },
    
    async cambiaUsername(){
      try
      {
        let response = await this.$axios.post('/user/' + this.profilo.ID + '/username', {username: this.newUsername}, { headers: {Authorization: "Bearer " + this.visitorID} })
        this.profilo.nome = this.newUsername;
        this.newUsername = "";
        localStorage.setItem('username', this.profilo.nome);
        alert("Username cambiato!");
        this.chiudiDialog();
      }
      catch(e)
      {
        this.errormsg = e.toString();
        alert(this.errormsg);
      }
    },
    async controllaSeguiti(){
      this.$router.push({path: '/userProfile/${this.profilo.ID}/following'})
    },
    async tornaHomePage(){
      localStorage.removeItem('IDCercato');
      this.$router.replace('/session');
    },
    async seguiProprietario(){
      try
      {
        let response = await this.$axios.post('/user/' + this.profilo.ID + '/username', {username: this.newUsername}, { headers: {Authorization: "Bearer " + this.visitorID} })
      }
      catch(e)
      {
        this.errormsg = e.toString();
        alert(this.errormsg);
      }
    }
  },
  mounted(){
    this.refresh();
  }
}
</script>

<template>
  <body>
    <header class=intestazione>
      <h4 @click = "tornaHomePage" >Torna alla HomePage</h4>
      <h1 style = "font-weight: bold;">Profilo di {{this.profilo.nome}}</h1>
      <div class=statistiche>
        <h4 @click = "controllaSeguiti">Followers: {{this.profilo.seguaci.length}}</h4>
        <h4>Seguiti: {{this.profilo.seguiti.length}}</h4>
      </div>
    </header>

    <dialog id = "cambiaNome" close>
      <h3 style = "text-align: center;"> Scrivi il tuo nuovo username </h3>
      <div style = "align-self: center;">
        <input placeholder="Nuovo username" v-model = this.newUsername style = "text-align: center;">
        <button @click = "cambiaUsername" :disabled= "this.newUsername.length > 16 || this.newUsername.length < 3">Cambia</button>
      </div>
    </dialog>

    <div class = riga>
      <div v-for = "immagine in this.profilo.fotoPubblicate">
        <div class = colonna>
          <img class = foto :src = immagine.File alt= immagine.photoID>
        </div>
      </div>
    </div>

    <div class = footer>
      
      <div class = ownerActions v-if = "this.profilo.ID == this.visitorID">
        <h4 style= "padding-left : 100px;"> Bannati: {{this.profilo.bannati.length}}</h4>
        <input id = "upload" type="file" accept="image/*" @change="uploadPhoto">
        <h4 style= "padding-right : 100px;" @click = "mostraDialog()"> Cambia username </h4>
      </div>

      <div class = visitorActions v-else>
        <button width = 200% style = "color:red;"> Banna </button>
        <button> Segui </button>
      </div>
    </div>
  </body>
</template>

<style>

  dialog[open]{
    position: absolute;
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
    position: fixed;
    width:100%;
    display: flex;
    align-items: center;
    justify-content:space-evenly;
    cursor: pointer;
    height:100px;

    background-color:brown;
    text-align: left;

    border-bottom-left-radius: 30px;
    border-bottom-right-radius: 30px;
    border-bottom: 3px solid black;
  }

  .riga{
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

    background-color:brown;

    border-top-left-radius: 30px;
    border-top-right-radius: 30px;
    border-top: 3px solid black;
  }

  .ownerActions{
    padding-top: 25px;
    display: flex;
    align-items: center;
    justify-content:space-evenly;
  }

  .visitorActions{
    padding-top: 25px;
    display: flex;
    justify-content: space-evenly;
  }
</style>