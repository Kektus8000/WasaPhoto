<script>
export default{
  data(){
    return{
      errormsg: "",
      profilo: {
        ID : localStorage.getItem('identifier'),
        username : localStorage.getItem('username'),
        seguaci: [],
        seguiti: [],
        bannati: [],
        fotoPubblicate: []
      },
      immagini: [{link:'https://www.avvenire.it/c/2017/PublishingImages/debda429421d455a8975d6ba03e67d65/caparezza.jpg?width=1024'},
                  {link: 'https://cdn-2.motorsport.com/images/amp/0ZRKlvo0/s1000/formula-1-spanish-gp-2023-char-2.jpg'},
                  {link: 'https://upload.wikimedia.org/wikipedia/it/2/22/Dragon_Ball_Super.png'},
                  {link: 'https://upload.wikimedia.org/wikipedia/it/2/22/Dragon_Ball_Super.png'},
                  {link:'https://www.avvenire.it/c/2017/PublishingImages/debda429421d455a8975d6ba03e67d65/caparezza.jpg?width=1024'},
                  {link: 'https://cdn-2.motorsport.com/images/amp/0ZRKlvo0/s1000/formula-1-spanish-gp-2023-char-2.jpg'}],
      commenti: ["Nice", "The Best", "Mille", "Buio", "No", "Ottimo"]
    }
  },
  methods:{
    async refresh(){
      this.recuperaInfo();
    },
    async visitaProfilo(checkID){
      try{
            let response = await this.$axios.get('/userProfile/' + checkID
            , 
              {
                headers: {Authorization: "Bearer " + this.profilo.ID}
              }
            );

            var risultato = response.data;

            localStorage.setItem('profiloCercato', risultato);
            this.$router.replace("/userProfile/" + checkID);
        }
        catch(e)
        {
            this.errormsg = e.toString();
            alert(this.errormsg);
        }
    },
    async recuperaInfo(){
      try
      {
        let response = await this.$axios.get('/userProfile/' + this.profilo.ID
            , 
              {
                headers: {Authorization: "Bearer " + this.profilo.ID}
              }
            );
        localStorage.setItem('profilo', response.data);
      }
      catch(e)
      {
        this.errormsg = e.toString();
        alert(this.errormsg);
      }
    },
    async logout(){
      localStorage.removeItem('profiloCercato');
      localStorage.removeItem('identifier');
      localStorage.removeItem('username');
      this.$router.replace("/");
    }
  },
    mounted() {
    this.refresh()
  } 
}
</script>

<template>
  <body>
    <div class = barraLaterale>
      <h2 class = introduzione>Benvenuto {{this.profilo.username}}</h2>
      <nav class = navigazione>
        <div class = opzioni style = "cursor: pointer">
          <h3 @click = "visitaProfilo(this.identifier)">Vai al tuo Profilo</h3>
          <h3>Seguiti</h3>
          <h3>Seguaci</h3>
          <h3 @click = "logout"> Logout </h3>
        </div>
      </nav>
    </div>

    <div class= contenuto>
      <template class=fotoPubblicate v-for = "foto in this.immagini">
        <div class = pubblicazione>
          <div class = foto>
            <div class = infoUser>
              <img class = miniPic width = 40 :src = foto.link>
              <h4 style = "padding-top: 5px; padding-left: 5px;">{{this.username}}</h4>
            </div>
            <img class = immagine :src = foto.link>
          </div>
          <div class = sezioneCommenti>
            <div v-for = "commento in this.commenti">
              <div class = commento>
                <div class = commentatore>
                  <img class = miniPic width = 30 height = 30 :src = foto.link>
                  <h5 style = "padding-top: 5px; padding-left: 5px;">{{this.username}}</h5>
                </div>
                <h4 style = "font-family: italic;">{{commento}}</h4>
              </div>
            </div>
          </div>
        </div>
      </template>
    </div>

  </body>

</template>

<style>
  .introduzione{
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

    background-color: orange;
    border-right: 5px solid black;
  }

  .contenuto{
    padding-left: 25%;
    width:100%;
  }

  .pubblicazione{
    height:400px;
    width:80%;
    display: flex;
    margin-bottom: 50px;
  }

  .infoUser{
    height: 40px;
    display:flex;
  }

  .miniPic{
    border-radius: 50%;
    border-style: solid;
    border-color: rgba(0,0,0,0.5);
  }
  
  .immagine{
    width:100%;
    height:90%;
  }

  .sezioneCommenti{
    top:40px;
    width:65%;
    display:inline;
    flex-wrap: wrap;
    overflow-y:scroll;
  }

  .commentatore{
    display: flex;
    text-align: center;
  }

  .commento{
    border: 1px solid black;
    word-wrap:break-word;
  }

</style>