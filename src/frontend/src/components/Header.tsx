import * as tsx from "vue-tsx-support";

export default tsx.component({
  name: "Header",
  render() {
    return (
      <header>
        <nav class="navbar navbar-light bg-light">
          <a class="navbar-brand">Topic Chat</a>

          <form class="form-inline">
            <input class="form-control mr-sm-2" type="search" placeholder="Search" aria-label="Search" />
            <button class="btn btn-outline-success my-2 my-sm-0" type="submit">Search</button>
          </form>
        </nav>
      </header>
    );
  }
});
