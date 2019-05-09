export default {
  name: 'MyMain',
  template: `
  <main>
        <div class="container">
            <ul class="center-text">
                <li><a href="">All</a></li>/
                <li><a href="">Design</a></li>/
                <li><a href="">Branding</a></li>/
                <li><a href="">Graphic</a></li>/
                <li><a href="">Animation</a></li>/
                <li><a href="">Illustration</a></li>/
                <li><a href="">Photography</a></li>
            </ul>
        </div>
        <div class="container">
            <div v-for="i in 15" class="col-3 padding-15">
                <div class="img-container">
                    <div class="img">
                        <a href=""><img src="images/anh2.png" alt=""></a>
                    </div>
                    <a href=""><img class="img-hover width_10" src="images/cart.png" alt=""></a>
                    <div class="content-text center-text">
                        <h3>Touch and Swipe</h3>
                        <h4>Branding/design</h4>
                    </div>
                </div>
            </div>
            

        </div>

        <div class="clearfix">
            
        </div>
    </main> 
  `
};
