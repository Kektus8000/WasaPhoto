<script>

const READER = new FileReader();

export default{
  data(){
    return{
      ricerca: "",
      errormsg: "",
      username : localStorage.getItem('username'),
      identifier: localStorage.getItem('identifier'),
      stream: [{link: ""}],
      utenti: []
    }
  },
  methods:{
    async refresh()
    {

    },
    async uploadPhoto(){
      try
      {
        let foto = document.getElementById('upload');
        READER.readAsArrayBuffer(foto.files[0]);
        READER.onload = async () => {
          let response = await this.$axios.put('/userProfile/${this.identifier}/publishedPhotos/', READER.result, { headers: {Authorization: "Bearer " + this.identifier} });
        };
      }
      catch(e)
      {
        this.errormsg = e.toString();
        alert(this.errormsg);
      }
    },
    async recuperaStream(){
      try
      {
        let response = await this.$axios.get('/userProfile/${this.identifier}/stream/', { headers: {Authorization: "Bearer " + this.identifier} });
        if (response.data != null) {for (let i = 0; i < response.data.length; i++){stream.push(response.data[i]);} }
      }
      catch(e)
      {
        this.errormsg = e.toString();
        alert(this.errormsg);
      }
    },
    async visitaProfilo(checkID){
        localStorage.setItem('IDCercato', checkID);
        this.$router.push({path: `/userProfile/${checkID}`});
    },
    logout(){
      localStorage.removeItem('IDCercato');
      localStorage.removeItem('identifier');
      localStorage.removeItem('username');
      this.$router.replace({path: "/"})
    }
  },
    watch: 
    {
      async ricerca(){
        if (this.ricerca != "")
        {
          try
          {
            let response = await this.$axios.put('/user/', {username: this.ricerca}, { headers: {Authorization: "Bearer " + this.identifier} });
            this.utenti = response.data;
            console.log(this.utenti);
          }
          catch(e)
          {
            this.errormsg = e.toString();
            alert(this.errormsg);
          }
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
    <div class = barraLaterale>
      <h2 class = introduzione>Benvenuto {{this.username}}</h2>
      <nav class = navigazione>
        <div class = opzioni style = "cursor: pointer">
          <input id = "upload" type="file" accept="image/*" @change="uploadPhoto">
          <input class = cercaNome placeholder ="Cerca Utente" v-model=this.ricerca>
          <h3 @click = "visitaProfilo(this.identifier)">Vai al tuo Profilo</h3>
          <h3 @click = "() => {this.$router.push({path: '/userProfile/${this.identifier}/following'}) }">Seguiti</h3>
          <h3>Seguaci</h3>
          <h3 @click = "logout"> Logout </h3>
        </div>
      </nav>
    </div>

    <div class= contenuto v-if = "this.ricerca == '' ">
      <div class=fotoPubblicate v-for = "foto in this.stream">
        <div class = pubblicazione>
          <div class = foto>
            <h4 height = 40px style = "font-weight: bold; padding-top: 5px; padding-left: 5px;">{{this.username}}</h4>
            <img class = immagine :src = foto.link>
          </div>
          <div class = sezioneCommenti>
            <div v-for = "commento in this.commenti">
              <div class = commento>
                <div class = commentatore>
                  <h5 style = "font-weight: bold; padding-top: 5px;">{{this.username}}</h5>
                </div>
                <h4 style = "font-family: italic; padding-left: 5px;">{{commento}}</h4>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class = risultato v-else>
      <template class= utenti-ricercati v-for = "user in this.utenti">
        <div class = utente-trovato>
          <h2 style = "padding-left:10px">{{user}}</h2>
        </div>
      </template>
    </div>

  </body>

</template>

<style>
  .introduzione{
    font-weight: bold;
    border-bottom: 5px solid black;
    word-wrap:break-word;
  }

  .barraLaterale{
    position: fixed;
    left:0%;
    top:0%;

    width: 15%;
    height: 100%;
    text-align: center;
    margin-top: 10px;

    background-color: orange;
    border-right: 5px solid black;
  }

  .cercaNome{
    width: 90%;
    height:30px;
    font-size: 25px;

    margin-bottom: 5px;
    cursor: pointer;
    align-self:right;
    border-radius: 5px;
    background-color: rgba(255, 255, 255, 0.5);
    
    color: white;
    font-family:'Gill Sans', 'Gill Sans MT', Calibri, 'Trebuchet MS', sans-serif;
    text-align: center;
  }

  .risultato{
    padding-left: 20%;
    width: 75%;
    height:100%;
  }

  .utente-trovato{
    padding-left: 10px;
    margin: 10px;

    display: flex;
    align-items: center;
    
    border: 1px solid black;
    box-shadow: 10px 10px 5px lightblue;
  }

  .contenuto{
    padding-left: 25%;
    width:100%;
    margin-bottom: 50px;
  }


  .pubblicazione{
    height:500px;
    width:80%;
    display: flex;
    margin-bottom: 50px;
    border: 1px solid black;
  }

  .immagine{
    width:100%;
    height: 455px;
  }


  .sezioneCommenti{
    top:40px;
    width:60%;
    height:100%;
    display:inline;
    flex-wrap: wrap;
    overflow-y:scroll;
  }

  .commento{
    border: 1px solid black;
    word-wrap:break-word;
  }

</style>