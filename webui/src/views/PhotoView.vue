<script>

const STATOLIKE = ["Mi Piace", "Ti Piace"];

export default{
  data(){
    return{
      errormsg: "",
      
      visitor : {
        visitorID : Number(localStorage.getItem('identifier')),
        visitorNome : localStorage.getItem('username'),
        visitorSeguaci : JSON.parse(localStorage.getItem('SeguaciSessione')),
        visitorSeguiti : JSON.parse(localStorage.getItem('SeguitiSessione')),
        visitorBannati : JSON.parse(localStorage.getItem('BannatiSessione'))
      },

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
        let temp = await this.$axios.get('/userProfile/' + this.profilo.ID + '/publishedPhotos/' + this.profilo.foto.PhotoID, 
        {responseType: 'blob',
        headers: {Authorization: "Bearer " + this.visitor.visitorID}});
        this.profilo.foto.Path = URL.createObjectURL(temp.data);

        this.profilo.foto.commentCount = this.profilo.foto.Comments != null ? this.profilo.foto.Comments.length : 0;
        this.profilo.foto.newComment = "";
    },
    async commentPhoto(photoID)
    {
      try
      {
        await this.$axios.put('/userProfile/' + this.profilo.ID + '/stream/' + photoID + '/comments/',
        { Text : this.profilo.foto.newComment},
        { headers: {Authorization: "Bearer " + this.visitor.visitorID } });
        this.profilo.foto.newComment = ""; 
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
        { headers: {Authorization: "Bearer " + this.visitor.visitorID } });
        this.refresh();
      }
      catch(e)
      {
        this.errormsg = e.toString();
        alert(this.errormsg);       
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
                <div height = 400px style = "border-bottom: 1px solid black;" v-for = "like in this.profilo.foto.Likes" :key = this.profilo.foto.Likes.UserID>
                    <h3 class= like style = "font-weight: bold; padding-top: 5px; cursor: pointer;" height = 40px
                    @click = visitaProfilo(like.UserID)> {{like.Username}} </h3>
                </div>
            </section>
            
            <section>
                <img class = foto-analizzata style= "border-radius: 10px;" :src = "this.profilo.foto.Path">
            </section>

            <section class = lista-commenti>
                <h3 height = 40px style = "top: 0; padding-top: 5px; border-bottom: 1px solid black;"> Commenti : {{this.profilo.foto.commentCount}} </h3>
                <div style = "border-bottom: 1px solid black;" v-for = "comm in this.profilo.foto.Comments" :key = comm.CommentID>
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