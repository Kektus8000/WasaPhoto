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
          <li >
            <span>
              <img class="Foto" :src= foto.link alt="Caricamento">
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
  
    .NavigazioneRapida{
      background-color: darkgoldenrod;
      text-shadow: 1px 2px 2px chocolate;
      color: black;
      border-block-end: solid;
      width: 100%;
      height:30px;
      display: flex;
      gap: 200px;
      margin-top: -10px;
      overflow-x: hidden;
      overflow-y: hidden;
        white-space: nowrap;
      align-items: center;
        -webkit-overflow-scrolling: touch;
    }
  
    .Foto{
      width:260px;
      height:260px;
      position:relative;
    }
  
    .ListaFoto{
      list-style:none;
      max-width: 700px;
    }
  
    .ListaFoto, li{
      display: flex;
      align-items: center;
      float:left;
      cursor:pointer;
      width:100%;
      gap: 1px;
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
                    {link: 'https://upload.wikimedia.org/wikipedia/it/2/22/Dragon_Ball_Super.png'}]
      }
    }
  }
  </script>