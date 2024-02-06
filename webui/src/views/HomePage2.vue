<template>
  <header class="Intestazione">
    <h3> Benvenuto {{this.username}}</h3>
    <nav class="NavigazioneRapida">
      <a href="link" target="_blank" style="text-decoration: none"> Followers </a>

      <span>
        <input placeholder="Cerca Utente" v-model="ricerca">
        <button :disabled="ricerca.length > 16 || ricerca.length < 3"> Cerca </button>
      </span>

      <a href="link" target="_blank" style="text-decoration: none"> Seguiti </a>
    </nav>
  </header>

  <main>
    <div class="ListaFoto">
      <ul v-for="foto in immagini">
        <li>
          <span>
            <img class="Foto" :src= foto.link alt="Caricamento">
            <button class="MiPiace"> Mi piace </button>
            <button class="Commenta"> Commenta </button>
          </span>
          
        </li>
      </ul>
    </div>

    <h1>{{this.ricerca}}</h1>
  </main>

</template>

<style>
  .Intestazione{
    text-align: center;
    text-shadow: 1px 2px 2px chocolate;
    background-color: darkgoldenrod;
    
  }

  .MiPiace{
    position:absolute;
    transform: translate(-350%, 420%);
    width:100px;
    height: 50px;
    background-color: gray;
    border-color: whitesmoke;
    border-radius: 20px;
  }

  .Commenta{
    position:absolute;
    transform: translate(-125%, 420%);
    width:100px;
    height: 50px;
    background-color: gray;
    border-color: whitesmoke;
    border-radius: 20px;
  }

  .NavigazioneRapida{
    background-color: darkgoldenrod;
    text-shadow: 1px 2px 2px chocolate;
    color: black;
    border-block-end: solid;
    width: 100%;
    height:30px;
    display: inline-flex;
    gap: 250px;
    margin-top: -10px;
    text-align: center;
    overflow-x: hidden;
    overflow-y: hidden;
    white-space: nowrap;
  }

  .Foto{
    width:380px;
    height:260px;
    border: 3px solid;
    border-color: black;
    box-shadow: 10px 10px;
  }

  .ListaFoto{
    display: flex;
    flex-wrap: wrap;
  }

  .ListaFoto, li{
    list-style: none;
    cursor:pointer;
    float:left;
    align-items: center;
  }
</style>

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
  }
}
</script>