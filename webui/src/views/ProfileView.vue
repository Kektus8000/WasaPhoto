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
        fotoPubblicate : [{
          photoID : 0,
          publisherID : 0,
          dataPubblicazione : null,
          path: "",
          commenti : [{
            commentID : 0,
            testo : "",
            photoID : 0,
            IDcommentatore : 0
          }]
        }]
      }
    }
  },
  methods: {
    async refresh(){
      await this.recuperaInfo();
      console.log(this.visitorID);
    },
    async recuperaInfo(){
      try
      {
        let response = await this.$axios.get('/userProfile/' + localStorage.getItem('IDCercato'), {headers: {Authorization: "Bearer " + this.visitorID}});
        this.profilo.nome = response.data.Username;
        if (response.data.Followers != null) {for (let i = 0; i < response.data.Followers.length; i++){this.profilo.seguaci.push(response.data.Followers[i]);} }
        if (response.data.Followings != null) {for (let i = 0; i < response.data.Followings.length; i++){this.profilo.seguiti.push(response.data.Followings[i]);} }
        if (response.data.Banneds != null) {for (let i = 0; i < response.data.Banneds.length; i++){bthis.profilo.bannati.push(response.data.Banneds[i]);} }
        if (response.data.PublishedPhotos != null) {
          var foto = response.data.PublishedPhotos;
          for (let i = 0; i < foto.length; i++){
            this.profilo.fotoPubblicate.push(foto[i]);
            var comms = foto[i].Comments;
            if (comms != null){ for (let j = 0; j < comms.length; j++) { this.profilo.fotoPubblicate[i].commenti.push(comms[j]); } }
            console.log(this.profilo.fotoPubblicate[i]);
          }
        }
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
      <h4 @click = "tornaHomePage" style = "padding-left: 10px;" >Torna alla HomePage</h4>
      <h1 style = "font-weight: bold;">Profilo di {{this.profilo.nome}}</h1>
      <div class=statistiche style = "padding-right: 10px;">
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
          <img class = foto :src = immagine.path alt= immagine.photoID>
        </div>
      </div>
    </div>

    <div class = footer>
        <h4 style= "padding-left:10px;" v-show = "this.profilo.ID == this.visitorID"> Bannati: {{this.profilo.bannati.length}}</h4>
        <input id = "upload" type="file" accept="image/*" @change="uploadPhoto" v-show = "this.profilo.ID == this.visitorID">
        <h4 style= "padding-right:10px;" v-show = "this.profilo.ID == this.visitorID" @click = "mostraDialog()"> Cambia username </h4>
    </div>
  </body>
</template>

<style>

  dialog[open]{
    position: absolute;
    top:  50%;
    left: 50%;
    display: inline;
    transform: translate(-50%, -10%);
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
    justify-content:space-between;
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
    border-radius: 10px;;
  }

  .footer{
    position:fixed;
    width:100%;
    display: flex;
    align-items: center;
    justify-content:space-between;
    cursor: pointer;
    height:100px;
    bottom: 0;

    background-color:brown;
    text-align: left;

    border-top-left-radius: 30px;
    border-top-right-radius: 30px;
    border-top: 3px solid black;
  }
</style>