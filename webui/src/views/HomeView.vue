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
      <h1 class = introduzione>Benvenuto {{this.username}}</h1>
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
          <img class = foto :src = foto.link>
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
    text-wrap:wrap;
  }

  .barraLaterale{
    position: fixed;
    left:0%;
    top:0%;

    width: 250px;
    height: 100%;
    text-align: center;

    background-color: orange;
    border-right: 5px solid black;
  }

  .contenuto{
    padding-left:500px;
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

</style>