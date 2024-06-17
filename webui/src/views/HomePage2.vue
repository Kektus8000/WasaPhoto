<script>
export default{
  data(){
    return{
      ricerca: "",
      username: "Simone",
      immagini: [{link:'https://www.avvenire.it/c/2017/PublishingImages/debda429421d455a8975d6ba03e67d65/caparezza.jpg?width=1024'},
                  {link: 'https://cdn-2.motorsport.com/images/amp/0ZRKlvo0/s1000/formula-1-spanish-gp-2023-char-2.jpg'},
                  {link: 'https://upload.wikimedia.org/wikipedia/it/2/22/Dragon_Ball_Super.png'},
                  {link: 'https://upload.wikimedia.org/wikipedia/it/2/22/Dragon_Ball_Super.png'},
                  {link:'https://www.avvenire.it/c/2017/PublishingImages/debda429421d455a8975d6ba03e67d65/caparezza.jpg?width=1024'},
                  {link: 'https://cdn-2.motorsport.com/images/amp/0ZRKlvo0/s1000/formula-1-spanish-gp-2023-char-2.jpg'}]
    }
  },
  methods:{
    
  }
}
</script>

<template>
  <body>
    <nav class = navigazione>
      <h1 class = introduzione>Benvenuto {{this.username}}</h1>
      <hr>
      <div class = opzioni>
        <RouterLink :to = " '/userProfile/' + '{{this.username}}' + '/following/'">
          <h3>Seguiti</h3>
        </RouterLink>
        <h3>Followers</h3>
      </div>
     </nav>


    <main class= schermata>
      <ul class=fotoPubblicate v-for = "foto in this.immagini">
        <img class = foto :src = foto.link>
       </ul>
    </main>
  </body>

</template>

<style>
  body{
    display:flex;
  }

  .navigazione{
    position: fixed;
    left:0%;
    top:0%;

    width: 225px;
    height: 100%;

    background-color: orange;
  }

  hr{color: black;}

  .opzioni{
    padding-left: 5px;
    text-align: left;
    cursor: pointer;
  }

  .fotoPubblicate{
    width: 100%;
    border-bottom: 3px solid black;
  }

  .foto{
    width: 50%;
  }
</style>