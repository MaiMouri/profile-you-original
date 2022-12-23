export async function getKeywords() {
    const headers = new Headers();
      headers.append("Content-Type", "application/json");
  
      const requestOptions = {
        method: "GET",
        headers: headers,
      }
    const res = await fetch('http://localhost:8080/keywords', requestOptions);
    const json = await res.json();
    if (!res.ok) throw new Error(json.message);
    return json;
  }


export async function createKeyword(keyword) {
    const requestBody = {
      Id: 0,
      Word: "",
      Description: "",
      ImageUrl: "",
      KeywordId: ""
    };
    requestBody.word = keyword.Word;
    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    // headers.append("Authorization", "Bearer " + jwtToken)
  
    const requestOptions = {
      method: "POST",
      headers: headers,
      body: JSON.stringify(requestBody),
    }
    const res = await fetch(`http://localhost:8080/keyword/create/${keyword.Word}`, requestOptions);
    const json = await res.json();
    if (!res.ok) throw new Error(json.message);
    // console.log(`keywordSlice received: ${json.Word}`);
    return json;
  }
  

export async function changeKeyword(keyword) {
    const requestBody = keyword;
    
    // passed validation, so save changes
    let headers = new Headers();
    headers.append("Content-Type", "application/json");

    const requestOptions = {
      method: "POST",
      headers: headers,
      body: JSON.stringify(requestBody)
    };
    
    const res = await fetch(`http://localhost:8080/keyword/update/`, requestOptions);
    const json = await res.json();
    if (!res.ok) throw new Error(json.message);
    console.log(json);
    return json;
  }

export async function removeKeyword(id) {
    let headers = new Headers();
    headers.append("Content-Type", "application/json");
    // headers.append("Authorization", "Bearer " + jwtToken)

    const requestBody = {
        Id: 0,
        Word: "",
        Description: "",
        ImageUrl: "",
        KeywordId: id,
      };

    const requestOptions = {
      method: "POST",
      headers: headers,
      body: JSON.stringify(requestBody)
    };
    
    const res = await fetch(`http://localhost:8080/keyword/delete/`, requestOptions);
    const json = await res.json();
    if (!res.ok) throw new Error(json.message);
    return json;
  }