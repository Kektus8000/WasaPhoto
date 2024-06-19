<script>
export default{
  data(){
    return{
      username: "Simone",
      seguiti: [{nome: "A", link:'https://www.avvenire.it/c/2017/PublishingImages/debda429421d455a8975d6ba03e67d65/caparezza.jpg?width=1024'},
                  {nome: "B", link: 'https://cdn-2.motorsport.com/images/amp/0ZRKlvo0/s1000/formula-1-spanish-gp-2023-char-2.jpg'},
                  {nome: "C", link: 'https://upload.wikimedia.org/wikipedia/it/2/22/Dragon_Ball_Super.png'},
                  {nome: "D", link: 'https://upload.wikimedia.org/wikipedia/it/2/22/Dragon_Ball_Super.png'},
                  {nome: "E", link:'https://www.avvenire.it/c/2017/PublishingImages/debda429421d455a8975d6ba03e67d65/caparezza.jpg?width=1024'},
                  {nome: "F", link: 'https://cdn-2.motorsport.com/images/amp/0ZRKlvo0/s1000/formula-1-spanish-gp-2023-char-2.jpg'},
                  {nome: "G", link:'https://www.avvenire.it/c/2017/PublishingImages/debda429421d455a8975d6ba03e67d65/caparezza.jpg?width=1024'},
                  {nome: "H", link: 'https://cdn-2.motorsport.com/images/amp/0ZRKlvo0/s1000/formula-1-spanish-gp-2023-char-2.jpg'},
                  {nome: "I", link: 'https://upload.wikimedia.org/wikipedia/it/2/22/Dragon_Ball_Super.png'},
                  {nome: "L", link: 'https://upload.wikimedia.org/wikipedia/it/2/22/Dragon_Ball_Super.png'},
                  {nome: "M", link:'https://www.avvenire.it/c/2017/PublishingImages/debda429421d455a8975d6ba03e67d65/caparezza.jpg?width=1024'},
                  {nome: "N", link: 'https://cdn-2.motorsport.com/images/amp/0ZRKlvo0/s1000/formula-1-spanish-gp-2023-char-2.jpg'}]
    }
  },
  methods:{
    stampa(){console.log("Simone");}
  }
}
</script>

<template>
  <body>
    <header class=intestazione>
      <h1>Account Seguiti</h1>
    </header>

    <div class = riga>
      <template class = listaFollowers v-for = "utente in this.seguiti">
        <div class = follower>
          <img class = icona :src = utente.link alt="avatar di {{utente.nome}}">
          <h2>{{utente.nome}}</h2>
          <button width = "15%">Vai al Profilo</button>
          <button class = blocca style = "color:red"  @click = stampa>Blocca</button>
        </div>
      </template>
    </div>
  </body>
</template>

<style>
  body{
    margin: 0;
    margin-right: 0;
    width:100%;

    display: inline-block;
    background-color: whitesmoke;
  }

  .intestazione{
    width:100%;
    height:100px;

    background-color: cadetblue;
    text-align: center;
  }

  .riga{
    display: flex;
    flex-wrap: wrap;
  }

  .follower{
    flex: 1 0 33.33%;
    text-align: center;
    border: 2px solid black;
  }

  .icona{
    border-radius: 50%;
    border-style: solid;
    border-color: rgba(0,0,0,0.5);

    width:95px;
    height:95px;
  }

</style>