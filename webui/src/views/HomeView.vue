<script>

const STATOLIKE = ["Mi Piace", "Ti Piace"]
const STATOUTENTE = ["Bannato", "Lo Segui", "Ti Segue"]

export default{
  data(){
    return{
      ricerca: "",

      errormsg: "",
      
      nuovoCommento : "",

      stream: [],
      
      utenti: [],

      profilo : {
        ID : Number(localStorage.getItem('identifier')),
        nome : localStorage.getItem('username'),
        seguaci : JSON.parse(localStorage.getItem('SeguaciSessione')),
        seguiti : JSON.parse(localStorage.getItem('SeguitiSessione')),
        bannati : JSON.parse(localStorage.getItem('BannatiSessione'))
      }
    }
  },
  methods:{
    async refresh()
    {
      await this.recuperaStream();
    },  
    async recuperaStream(){
      try
      {
        let response = await this.$axios.get('/userProfile/' + this.profilo.ID + '/stream/', { headers: {Authorization: "Bearer " + this.profilo.ID} });
        if (response.data != null) {
          this.stream = response.data;
          for (let i = 0; i < this.stream.length; i++)
          {
            var foto = this.stream[i];
            let temp = await this.$axios.get('/userProfile/' + foto.PublisherID + '/publishedPhotos/' + foto.PhotoID, 
            {responseType: 'blob',
            headers: {Authorization: "Bearer " + this.profilo.ID} });
            foto.File = URL.createObjectURL(temp.data);
            
            let temp2 = await this.$axios.get('/user/' + foto.PublisherID, {});
            foto.PublisherName = temp2.data.Username;

            foto.isLiked = false;
            if (foto.Likes != null) {foto.isLiked = foto.Likes.some(like => like.UserID === this.profilo.ID);}
            foto.initiallyLiked = foto.isLiked;
            foto.likeCount = foto.Likes != null ? foto.Likes.length : 0;

            foto.commentCount = foto.Comments != null ? foto.Comments.length : 0;
            foto.newComment = '';
            console.log(foto.Comments);
          }  
        }
      }
      catch(e)
      {
        this.errormsg = e.toString();
        alert(this.errormsg);
      }
    },
    fotoPiaciuta(photoID)
    {
      var foto = this.stream.find(foto => foto.PhotoID === photoID);
      return foto.isLiked ? STATOLIKE[1] : STATOLIKE[0]
    },
    likePhoto(photoID)
    {
      var foto = this.stream.find(foto => foto.PhotoID === photoID);
      foto.isLiked = !foto.isLiked;
      if (foto.isLiked) {foto.likeCount += 1}
      else {foto.likeCount -= 1}
    },
    async salvaMiPiace()
    {
      try
      {
        for (let i = 0; i < this.stream.length; i++)
        {
          var foto = this.stream[i];
          if (foto.isLiked !== foto.initiallyLiked)
          {
            if (foto.isLiked) { await this.$axios.post('/userProfile/' + this.profilo.ID + '/stream/' + foto.PhotoID + '/likes/', {}, { headers: {Authorization: "Bearer " + this.profilo.ID} });}
            else {await this.$axios.delete('/userProfile/' + this.profilo.ID + '/stream/' + foto.PhotoID + '/likes/', { headers: {Authorization: "Bearer " + this.profilo.ID} });}
          }
        }
      }
      catch(e)
      {
        this.errormsg = e.toString();
        alert(this.errormsg);       
      }
    },
    async pubblicaCommento(photoID)
    {
      try
      {
        var foto = this.stream.find(foto => Number(foto.PhotoID) === Number(photoID));
        await this.$axios.put('/userProfile/' + this.profilo.ID + '/stream/' + photoID + '/comments/',
        { Text : foto.newComment},
        { headers: {Authorization: "Bearer " + this.profilo.ID} });
        foto.newComment = ""; 
        this.salvaMiPiace();
        this.refresh();
      }
      catch(e)
      {
        this.errormsg = e.toString();
        alert(this.errormsg);       
      }
    },
    async uncommentPhoto(photoID, commentID)
    {
      try
      {
        await this.$axios.delete('/userProfile/' + this.profilo.ID + '/stream/' + photoID + '/comments/' + commentID, 
        { headers: {Authorization: "Bearer " + this.profilo.ID} });
        this.salvaMiPiace();
        this.refresh();
      }
      catch(e)
      {
        this.errormsg = e.toString();
        alert(this.errormsg);       
      }
    },
    /////////////////////////////////////// METODI PER RAGGIUNGERE ALTRE PAGINE /////////////////////////////////////
    async visitaProfilo(checkID){
      await this.salvaMiPiace();
      localStorage.setItem('IDCercato', checkID);
      this.$router.push({path: '/userProfile/' + checkID });
    },
    async visitaSeguiti(){
      await this.salvaMiPiace();
      localStorage.setItem('IDCercato', this.profilo.ID);
      this.$router.push({path: '/userProfile/' + this.profilo.ID + '/followeds'});
    },
    async visitaSeguaci(){
      await this.salvaMiPiace();
      localStorage.setItem('IDCercato', this.profilo.ID);
      this.$router.push({path: '/userProfile/' + this.profilo.ID + '/following'});
    },
    async logout(){
      await this.salvaMiPiace();
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
  <head>
    <link rel="stylesheet" href="https://fonts.googleapis.com/icon?family=Material+Icons">
  </head>
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
      <div v-for = "foto in this.stream" :key = foto.PhotoID>
        <div class = pubblicazione>
          <section class = sezione-foto>
            <h4 height = 40px style = "font-weight: bold; padding-top: 5px; padding-left: 5px; cursor: pointer;" @click = visitaProfilo(foto.PublisherID)> {{foto.PublisherName}}</h4>
            <img class = immagine :src = foto.File>
            <button class = like-stream  @click = likePhoto(foto.PhotoID)> <i class="material-icons">favorite</i> {{fotoPiaciuta(foto.PhotoID)}} : {{foto.likeCount}} </button>
          </section>
          <section class = sezioneCommenti>
            <h3 height = 40px style = "top: 0; padding-top: 5px; border-bottom: 1px solid black;"> Commenti : {{foto.commentCount}} </h3>
            <div height = 400px v-for = "comm in foto.Comments" :key = comm.CommentID>
              <div class = commento>
                <div class = commentatore style = "display: flex; justify-content: space-between">
                  <h5 style = "font-weight: bold; padding-top: 5px;"> {{comm.PublisherName}}</h5>
                  <button v-if = "Number(comm.PublisherID) === Number(this.profilo.ID)"
                  @click = "uncommentPhoto(foto.PhotoID, comm.CommentID)" style = "background-color: rgb(178, 34, 34); color: white;"> Cancella </button>
                </div>
                <h4 style = "font-family: italic; padding-left: 5px;">{{comm.Text}}</h4>
              </div>
            </div>
            <div class = aggiungi-commento height = 40px>
              <input v-model = foto.newComment height = 40px width = 400>
              <button @click = pubblicaCommento(foto.PhotoID)> Pubblica </button>
            </div>
          </section>
        </div>
      </div>
    </div>

    <div class = risultato v-else>
      <div class= utenti-ricercati v-for = "user in this.utenti" :key = user.UserID>
        <div class = utente-trovato>
          <h3 style = "font-weight: bold; padding-left:10px" @click = "visitaProfilo(user.UserID)"> {{user.Username}}</h3>
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
    width:90%;
    display: flex;
    margin-bottom: 50px;
    border: 1px solid black;
  }

  .sezione-foto{
    width: 50%;
  }
  .immagine{
    width: 100%;
    height: 455px;
  }

  .sezioneCommenti{
    border-left: 1px solid black;
    width:50%;
    height:100%;
    display:inline;
    overflow-y:scroll;
  }

  .commento{
    border: 1px solid black;
    word-wrap:break-word;
  }

  .like-stream{
    transform: translate(10%, -120%);
    position: relative;

    font-size: 20px;

    border-radius:10px;
    border: 1px solid black;
    color: rgb(80, 200, 120);
    cursor: pointer;
    background-color: white;
  }

  .like-stream:hover{
    color: white;
    background-color: rgb(80, 200, 120);
  }

  .aggiungi-commento{
    border-top: 1px solid black;
    bottom: 0;
    width: 100%;
  }

</style>