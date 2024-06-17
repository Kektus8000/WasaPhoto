<script>
export default{
  data(){
    return{
      ricerca: "",
      username: "Simone",
      seguiti: [{nome: "Simone", link:'https://www.avvenire.it/c/2017/PublishingImages/debda429421d455a8975d6ba03e67d65/caparezza.jpg?width=1024'},
                  {nome: "Marta", link: 'https://cdn-2.motorsport.com/images/amp/0ZRKlvo0/s1000/formula-1-spanish-gp-2023-char-2.jpg'},
                  {nome: "Ansh", link: 'https://upload.wikimedia.org/wikipedia/it/2/22/Dragon_Ball_Super.png'},
                  {nome: "Nicolas", link: 'https://upload.wikimedia.org/wikipedia/it/2/22/Dragon_Ball_Super.png'},
                  {nome: "Lavinia", link:'https://www.avvenire.it/c/2017/PublishingImages/debda429421d455a8975d6ba03e67d65/caparezza.jpg?width=1024'},
                  {nome: "Edoardo", link: 'https://cdn-2.motorsport.com/images/amp/0ZRKlvo0/s1000/formula-1-spanish-gp-2023-char-2.jpg'},
                  {nome: "Simone", link:'https://www.avvenire.it/c/2017/PublishingImages/debda429421d455a8975d6ba03e67d65/caparezza.jpg?width=1024'},
                  {nome: "Marta", link: 'https://cdn-2.motorsport.com/images/amp/0ZRKlvo0/s1000/formula-1-spanish-gp-2023-char-2.jpg'},
                  {nome: "Ansh", link: 'https://upload.wikimedia.org/wikipedia/it/2/22/Dragon_Ball_Super.png'},
                  {nome: "Nicolas", link: 'https://upload.wikimedia.org/wikipedia/it/2/22/Dragon_Ball_Super.png'},
                  {nome: "Lavinia", link:'https://www.avvenire.it/c/2017/PublishingImages/debda429421d455a8975d6ba03e67d65/caparezza.jpg?width=1024'},
                  {nome: "Edoardo", link: 'https://cdn-2.motorsport.com/images/amp/0ZRKlvo0/s1000/formula-1-spanish-gp-2023-char-2.jpg'}]
    }
  },
  methods:{
    
  }
}
</script>

<template>
  <body>
    <header class=intestazione>
      <h1>Account Seguiti</h1>
    </header>
  
    <template class = listaFollowers v-for = "utente in this.seguiti">
      <div class = follower>
        <span class = info>
          <img class = icona :src = utente.link alt="avatar di {{utente.nome}}">
          <h2>{{utente.nome}}</h2>
        </span>
        <span class = opzioni>
          <h4>Vai al Profilo</h4>
          <h4 class = blocca>Blocca</h4>
          <div class=vl></div>
        </span>
      </div>
    </template>
  </body>
</template>

<style>
  body{
    margin: 0;
    margin-right: 0;
    width:100%;
    margin-top: 0;

    background-color: whitesmoke;
  }
  .intestazione{
    width:100%;
    height:100px;


    background-color: cadetblue;
    text-align: center;
  }

  .follower{
    max-width:25%;
    align-items: center;
    text-align: center;
    display: inline-block;
    justify-content: space-evenly;
    border: 2px solid black;
  }

  .icona{
    /* Bordo dell'icona */
    border-radius: 50%;
    border-style: solid;
    border-color: rgba(0,0,0,0.5);

    /* Posizione e allineamento*/
    width:100px;
    height:100px;
  }
  
  h4{
    cursor: pointer;
  }

  .blocca{
    color: red;
  }
</style>