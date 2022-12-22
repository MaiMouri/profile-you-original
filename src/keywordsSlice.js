import { createSlice } from '@reduxjs/toolkit';

import { KeywordsData } from "./DummyData";

async function getKeywords() {
  const headers = new Headers();
    headers.append("Content-Type", "application/json");

    const requestOptions = {
      method: "GET",
      headers: headers,
    }
  const res = await fetch('http://localhost:8080/keywords', requestOptions);
  const json = await res.json();
  if (!res.ok) throw new Error(json.message);
  console.log(json);
  return json;
}

async function createKeyword(keyword) {
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

  const requestOptions = {
    method: "POST",
    headers: headers,
    body: JSON.stringify(requestBody),
  }
  const res = await fetch(`http://localhost:8080/keyword/create/${keyword.Word}`, requestOptions);
  const json = await res.json();
  if (!res.ok) throw new Error(json.message);
  console.log(json);
  return json;
}


//createSliceを使ったら reducer を作成するだけで自動的に action type も定義してくれるし action creator も生成。
export const keywordSlice = createSlice({
  //name -> action type の prefix
  name: "items",
  initialState: { loadingNow: false, error: null, items: [] },
  // initialState: { value: KeywordsData },
  reducers: {
    //addUser -> action type
    addKeyword: (state, action) => {
      console.log(state);

      state.items = [action.payload, ...state.items]
    },

    deleteKeyword: (state, action) => {
      state.items = state.items.filter((item) => item.KeywordId !== action.payload.KeywordId);
    },
     // 通信を開始した時に呼ぶ関数
     fetchStart(state, action) {
      state.loadingNow = true;
      state.error = null;
    },
    // 通信が失敗した時に呼ぶ関数
    fetchFailure(state, action) {
      state.loading = false;
      state.error = action.payload;
    },
    // 通信が成功した時に呼ぶ関数
    fetchSuccess(state, action) {
      state.loadingNow = false;
      state.error = null;
      state.items = action.payload;
    },
  },
});

// Actions
export const { addKeyword, deleteKeyword, fetchStart, fetchFailure, fetchSuccess } = keywordSlice.actions;


// 外部からはこの関数を呼んでもらう
export const fetchItems = () => async dispatch => {
  try {
    dispatch(fetchStart());
    dispatch(fetchSuccess(await getKeywords()));
  } catch (error) {
    dispatch(fetchFailure(error.stack));
  }
};
export const postKeyword = (word) => async dispatch => {
  try {
    dispatch(addKeyword(await createKeyword(word)));
  } catch (error) {
    dispatch(fetchFailure(error.stack));
  }
};

// Selectors keywords-storeのreducer nameと揃える
export const selectKeyword = ({ keywords }) => keywords;

// Reducer(must be default export)
export default keywordSlice.reducer;