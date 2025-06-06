import Component from "@glimmer/component";
import { tracked } from "@glimmer/tracking";
import { action } from "@ember/object";

export default class SearchComponent extends Component {
  @tracked subject_code;

  get subject_code() {
    return this.subject_code;
  }

  @action
  async search() {
    let params = new URLSearchParams({
      "subject_code": this.subject_code
    });

    let resp = await fetch(
      "http://localhost:42069/search?" + params.toString(),
      {method: "POST"}
    );

    if (resp.status !== 200) {
      alert("Something went wrong");
      return;
    }

    let json = await resp.json()
    console.log(json);
  }
}
