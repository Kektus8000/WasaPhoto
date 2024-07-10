<script>

const STATOLIKE = ["Mi Piace", "Ti Piace"];

export default{
  data(){
    return{
      visitor : {
        visitorID : Number(localStorage.getItem('identifier')),
        visitorNome : localStorage.getItem('username'),
        visitorSeguaci : JSON.parse(localStorage.getItem('SeguaciSessione')),
        visitorSeguiti : JSON.parse(localStorage.getItem('SeguitiSessione')),
        visitorBannati : JSON.parse(localStorage.getItem('BannatiSessione'))
      },

      comments: [],

      profilo : {
        ID : Number(localStorage.getItem('IDCercato')),
        foto: JSON.parse(localStorage.getItem('FotoAnalizzata'))
      }
    }
  },
  methods: {
    async refresh(){
        this.retrievePhoto();
    },
    async retrievePhoto(){
      try
      {
        let temp = await this.$axios.get('/userProfile/' + this.profilo.ID + '/publishedPhotos/' + this.profilo.foto.PhotoID, 
        {responseType: 'blob',
        headers: {Authorization: "Bearer " + this.visitor.visitorID}});
        this.profilo.foto.Path = URL.createObjectURL(temp.data);

        this.profilo.foto.commentCount = this.profilo.foto.Comments != null ? this.profilo.foto.Comments.length : 0;
        this.profilo.foto.newComment = "";
        if (this.profilo.foto.Comments != null) {this.comments = this.profilo.foto.Comments;}
      }
      catch(e)
      {
        if (e.response != null)
          {
            switch(e.response.status)
            {
              case 500 : 
                alert("Un errore nel server impedisce l'operazione!");
                break;
            }
          }
      }
    },
    async commentPhoto(photoID)
    {
      try
      {
        let response = await this.$axios.post('/userProfile/' + this.profilo.ID + '/stream/' + photoID + '/comments/',
        { Text : this.profilo.foto.newComment},
        { headers: {Authorization: "Bearer " + this.visitor.visitorID } });
        
        var commento = {CommentID : response.data,
        Text : this.profilo.foto.newComment,
        PhotoID: this.profilo.foto.PhotoID,
        PublisherID : this.visitor.visitorID, 
        PublisherName: this.visitor.visitorNome};
        this.profilo.foto.newComment = "";
        
        if (this.comments == null) {this.comments = commento;}
        else {this.comments.push(commento);}
        this.profilo.foto.commentCount += 1;
      }
      catch(e)
      {
        if (e.response != null)
        {
          switch (e.response.status)
          {
            case 400:
              alert("La lunghezza del commento non è corretta!\n(Minimo 6 caratteri, massimo 160 caratteri)");
              break;
            case 500:
              alert("Un errore nel server impedisce l'operazione!");
              break;
          }
        }  
      }
    },
    async uncommentPhoto(photoID, commentID)
    {
      try
      {
        await this.$axios.delete('/userProfile/' + this.profilo.ID + '/stream/' + photoID + '/comments/' + commentID, 
        { headers: {Authorization: "Bearer " + this.visitor.visitorID } });
        let indice = this.comments.map(comm => comm.CommentID).indexOf(commentID);
        this.comments.splice(indice, 1);
        
        this.profilo.foto.commentCount -= 1;
      }
      catch(e)
      {
        if (e.response != null)
        {
          switch (e.response.status)
          {
            case 403:
              alert("Non puoi rimuovere un commento se non è tuo o se non sei il proprietario della foto!");
              break;
            case 404:
              alert("Errore nell'autenticazione!");
              break;
            case 500:
              alert("Un errore nel server impedisce l'operazione!");
              break;
          }
        }        
      }
    },
    async tornaIndietro(){
      localStorage.removeItem('FotoAnalizzata');
      this.$router.back();
    },
    async visitaProfilo(checkID){
      localStorage.removeItem('FotoAnalizzata');
      localStorage.setItem('IDCercato', checkID);
      this.$router.push({path: '/userProfile/' + checkID });
    },
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
        <header class=intestazione>
            <h3 @click = tornaIndietro() style= "padding-left : 10px; font-weight: bold;">Torna Indietro</h3>
        </header>

        <div height = 400px class = statistiche-foto>
            <section class = sezione-like>
                <h4 height = 40px style = "top: 0; padding-top: 5px; border-bottom: 1px solid black;"> Like Ricevuti : {{this.profilo.foto.Likes ? this.profilo.foto.Likes.length : 0}} </h4>
                <div height = 400px style = "border-bottom: 1px solid black;" v-for = "like in this.profilo.foto.Likes" :key = like.UserID>
                    <h3 class= like style = "font-weight: bold; padding-top: 5px; cursor: pointer;" height = 40px
                    @click = visitaProfilo(like.UserID)> {{like.Username}} </h3>
                </div>
            </section>
            
            <section>
                <img class = foto-analizzata style= "border-radius: 10px;" :src = "this.profilo.foto.Path">
            </section>

            <section class = lista-commenti>
                <h3 height = 40px style = "top: 0; padding-top: 5px; border-bottom: 1px solid black;"> Commenti : {{this.profilo.foto.commentCount}} </h3>
                <div style = "border-bottom: 1px solid black;" v-for = "comm in this.comments" :key = comm.CommentID>
                    <div>
                        <div class = commentatore style = "display: flex; justify-content: space-between">
                            <h5 style = "font-weight: bold; padding-top: 5px; cursor: pointer;"
                            @click = visitaProfilo(comm.PublisherID)> {{comm.PublisherName}}</h5>
                            <button v-if = "Number(comm.PublisherID) === Number(this.visitor.visitorID) || Number(this.visitor.visitorID) === Number(this.profilo.ID)"
                            @click = "uncommentPhoto(this.profilo.foto.PhotoID, comm.CommentID)" style = "background-color: rgb(178, 34, 34); color: white;"> Cancella </button>
                        </div>
                        <h4 style = "font-family: italic; padding-left: 5px;">{{comm.Text}}</h4>
                    </div>
                </div>
                <div class = aggiungi-commento height = 40px>
                    <input v-model = this.profilo.foto.newComment placeholder = "Scrivi un commento" size = 55>
                    <button @click = commentPhoto(this.profilo.foto.PhotoID)> Pubblica </button>
                </div>
            </section>
        </div>
    </body>
</template>

<style>
    .statistiche-foto{
        display: flex;
        padding-top: 150px;
        margin-left: auto;
        margin-right: auto;
        width: 100%;
    }
    .foto-analizzata{
        flex: 33.33%;
        padding-left: 5px;
        padding-right: 5px;
        float:left;
        width:100%;
        height: 400px;
        align-self: center;
    }

    .sezione-like{
        border: 1px solid black;
        flex: 33.33%;
        float:left;
        height:400px;
        display:inline;
        overflow-y:scroll;
    }
    .lista-commenti{
        height: 400px;
        float:left;
        flex: 33.33%;
        border: 1px solid black;
        display:inline;
        overflow-y:scroll;
    }


</style>