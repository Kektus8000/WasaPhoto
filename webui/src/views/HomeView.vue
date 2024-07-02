<script>

export default{
  data(){
    return{
      ricerca: "",
      errormsg: "",
      
      stream: [],
      
      utenti: [],

      profilo : {
        ID : Number(localStorage.getItem('identifier')),
        nome : localStorage.getItem('username'),
        seguaci : [],
        seguiti : [],
        bannati : [],
      }
    }
  },
  methods:{
    async refresh()
    {
      await this.recuperaInfo();
      await this.recuperaStream();
      console.log(this.profilo.ID);
      console.log(this.profilo.nome);
    },    
    async recuperaInfo(){
      try
      {
        let response = await this.$axios.get('/userProfile/' + this.profilo.ID, {headers: {Authorization: "Bearer " + this.profilo.ID}});
        if (response.data.Followers != null) {this.profilo.seguaci = response.data.Followers; }
        if (response.data.Followings != null) {this.profilo.seguiti = response.data.Followings; }
        if (response.data.Banneds != null) {this.profilo.bannati = response.data.Banneds;}
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
        let response = await this.$axios.get('/userProfile/${this.profilo.ID}/stream/', { headers: {Authorization: "Bearer " + this.profilo.ID} });
        if (response.data != null) { this.stream = response.data; }
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
    async visitaSeguiti(){
      localStorage.setItem('IDCercato', this.profilo.ID);
      this.$router.push({path: '/userProfile/' + this.profilo.ID + '/followeds'});
    },
    async visitaSeguaci(){
      localStorage.setItem('IDCercato', this.profilo.ID);
      this.$router.push({path: '/userProfile/' + this.profilo.ID + '/following'});
    },
    async logout(){
      localStorage.removeItem('IDCercato');
      localStorage.removeItem('identifier');
      localStorage.removeItem('username');
      this.$router.replace({path: "/"});
    },
  },
    watch: 
    {
      async ricerca(){
        if (this.ricerca != "")
        {
          try
          {
            let response = await this.$axios.put('/user/', {username: this.ricerca}, { headers: {Authorization: "Bearer " + this.profilo.ID} });
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
      <h2 class = introduzione>Benvenuto {{this.profilo.nome}}</h2>
      <nav class = navigazione>
        <div class = opzioni style = "cursor: pointer">
          <input class = cercaNome placeholder ="Cerca Utente" v-model=this.ricerca>
          <h3 @click = "visitaProfilo(this.profilo.ID)">Vai al tuo Profilo</h3>
          <h3 @click = "visitaSeguiti">Seguiti</h3>
          <h3 @click = "visitaSeguaci">Seguaci</h3>
          <h3 @click = "logout"> Logout </h3>
        </div>
      </nav>
    </div>

    <div class= contenuto v-if = "this.ricerca == '' ">
      <div class=fotoPubblicate v-for = "foto in this.stream">
        <div class = pubblicazione>
          <div class = foto>
            <h4 height = 40px style = "font-weight: bold; padding-top: 5px; padding-left: 5px;">{{this.profilo.nome}}</h4>
            <img class = immagine :src = foto.link>
          </div>
          <div class = sezioneCommenti>
            <div v-for = "commento in this.commenti">
              <div class = commento>
                <div class = commentatore>
                  <h5 style = "font-weight: bold; padding-top: 5px;">{{this.profilo.nome}}</h5>
                </div>
                <h4 style = "font-family: italic; padding-left: 5px;">{{commento}}</h4>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class = risultato v-else>
      <div class= utenti-ricercati v-for = "user in this.utenti">
        <div class = utente-trovato>
          <h2 style = "font-weight: bold; padding-left:10px" @click = "visitaProfilo(user.UserID)"> {{user.Username}} </h2>
          <h2 style = "color: green; padding-top: 5px; padding-right: 20px;" v-if = "this.profilo.seguiti.includes(user.userID)"> Seguito </h2>
          <h2 style = "color: red; padding-top: 5px; padding-right: 20px;" v-if = "this.profilo.bannati.includes(user.userID)"> Bannato </h2>
        </div>
      </div>
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

    cursor: pointer;
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