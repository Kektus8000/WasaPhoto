<script>
export default{
  data(){
    return{
      username: localStorage.getItem('username'),
      immagini: [{link:'https://www.avvenire.it/c/2017/PublishingImages/debda429421d455a8975d6ba03e67d65/caparezza.jpg?width=1024'},
                  {link: 'https://cdn-2.motorsport.com/images/amp/0ZRKlvo0/s1000/formula-1-spanish-gp-2023-char-2.jpg'},
                  {link: 'https://upload.wikimedia.org/wikipedia/it/2/22/Dragon_Ball_Super.png'},
                  {link: 'https://upload.wikimedia.org/wikipedia/it/2/22/Dragon_Ball_Super.png'},
                  {link:'https://www.avvenire.it/c/2017/PublishingImages/debda429421d455a8975d6ba03e67d65/caparezza.jpg?width=1024'},
                  {link: 'https://cdn-2.motorsport.com/images/amp/0ZRKlvo0/s1000/formula-1-spanish-gp-2023-char-2.jpg'}],
      commenti: ["Simoneeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee", "The Best", "Ansh", "Gayest", "No", "Gay"]
    }
  },
  methods:{
    
  }
}
</script>

<template>
  <body>
    <div class = barraLaterale>
      <h2 class = introduzione>Benvenuto {{this.username}}</h2>
      <nav class = navigazione>
        <div class = opzioni style = "cursor: pointer">
          <RouterLink :to = "'/userProfile/' + this.username" style="text-decoration: none; color: inherit;">
            <h3>Vai al tuo Profilo</h3>
          </RouterLink>
          
          <RouterLink :to = "'/'" style="text-decoration: none; color: inherit;">
            <h3>Logout</h3>
          </RouterLink>
        </div>
      </nav>
    </div>

    <div class= contenuto>
      <template class=fotoPubblicate v-for = "foto in this.immagini">
        <div class = pubblicazione>
          <div class = foto>
            <div class = infoUser>
              <img class = icona :src = foto.link>
              <h5>{{this.username}}</h5>
            </div>
            <img class = immagine :src = foto.link>
          </div>
          <div class = sezioneCommenti>
            <div v-for = "commento in this.commenti">
              <div class = commento>
                <h5 style = "font-family: italic;">{{this.username}}</h5>
                <h3>{{commento}}</h3>
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
    border: 1px solid black;
  }

  .sezioneCommenti{
    padding-left: 5px;
    width:60%;
    display:inline;
    flex-wrap: wrap;
    overflow-y:scroll;
  }

  .commento{
    border: 1px solid black;
    word-wrap:break-word;
  }

  .infoUser{
    height: 10%;
    display:flex;
  }

  .immagine{
    width:100%;
    height:90%;
  }

  .icona{
    flex: 0.25;
    border-radius: 50%;
    border-style: solid;
    border-color: rgba(0,0,0,0.5);

    width:10px;
    height:10px;
  }

</style>