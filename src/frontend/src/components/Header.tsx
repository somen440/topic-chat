import * as tsx from "vue-tsx-support";
import * as user from "@/store/modules/user";

export default tsx.component({
  name: "Header",
  render() {
    return (
      <header>
        <div class="bg-info collapse" id="navbarHeader" style="">
          <div class="container">
            <div class="row">
              <div class="col-sm-8 col-md-7 py-4">
                <h4 class="text-white">About</h4>
                <p class="text-muted">todo サービス補足的なあれをいれるよ</p>
              </div>
              <div class="col-sm-4 offset-md-1 py-4">
                <h4 class="text-white">運営者情報</h4>
                <ul class="list-unstyled">
                  <li>
                    <a href="#" class="text-white">
                      todo Follow on Twitter
                    </a>
                  </li>
                  <li>
                    <a href="#" class="text-white">
                      todo Like on Facebook
                    </a>
                  </li>
                  <li>
                    <a href="#" class="text-white">
                      todo Email me
                    </a>
                  </li>
                </ul>
              </div>
            </div>
          </div>
        </div>

        <div class="navbar navbar-dark bg-info shadow-sm">
          <div class="container d-flex justify-content-between">
            <a
              href="#"
              class="navbar-brand d-flex align-items-center text-white"
            >
              <strong>Topic Chat</strong>
            </a>

            {user.readIsLoggedIn(this.$store) ? (
              <form class="form-inline">
                <input
                  class="form-control mr-sm-2"
                  type="search"
                  placeholder="Search"
                  aria-label="Search"
                />
                <button
                  class="btn btn-outline-success my-2 my-sm-0"
                  type="submit"
                >
                  Search
                </button>
              </form>
            ) : (
              <button
                class="navbar-toggler collapsed"
                type="button"
                data-toggle="collapse"
                data-target="#navbarHeader"
                aria-controls="navbarHeader"
                aria-expanded="false"
                aria-label="Toggle navigation"
              >
                <span class="navbar-toggler-icon"></span>
              </button>
            )}
          </div>
        </div>
      </header>
    );
  }
});
