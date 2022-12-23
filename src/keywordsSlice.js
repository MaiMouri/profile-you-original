import { createSlice } from '@reduxjs/toolkit';
import { getKeywords, createKeyword, changeKeyword, removeKeyword } from './api';
import { KeywordsData } from "./DummyData";


//createSliceを使ったら reducer を作成するだけで自動的に action type も定義してくれるし action creator も生成。
export const keywordSlice = createSlice({
  //name -> action type の prefix
  name: "items",
  initialState: { loadingNow: false, error: null, items: [] },
  // initialState: { value: KeywordsData },
  reducers: {
    //addUser -> action type
    addKeyword: (state, action) => {
      state.items = [action.payload, ...state.items]
    },
    updateKeyword: (state, action) => {
      const { KeywordId, Word, ImageUrl, Description } = action.payload
      const existingKeyword = state.items.find((keyword) => keyword.KeywordId === KeywordId);
      console.log("upadtaKeyword in slice" + existingKeyword);
      console.log(KeywordId, Word, ImageUrl, Description);
      // state.items = state.items.map((item) => item.KeywordId === action.payload.KeywordId);
      if (existingKeyword) {
        existingKeyword.Word = Word;
        existingKeyword.Description = Description;
        existingKeyword.ImageUrl = ImageUrl;
      }
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
export const { addKeyword, updateKeyword, deleteKeyword, fetchStart, fetchFailure, fetchSuccess } = keywordSlice.actions;


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

export const renewalKeyword = (id) => async dispatch => {
  try {
    dispatch(updateKeyword(await changeKeyword(id)));
  } catch (error) {
    dispatch(fetchFailure(error.stack));
  }
};

export const delKeyword = (id) => async dispatch => {
  try {
    dispatch(deleteKeyword(await removeKeyword(id)));
  } catch (error) {
    dispatch(fetchFailure(error.stack));
  }
};

// Selectors keywords-storeのreducer nameと揃える
export const selectKeyword = ({ keywords }) => keywords;

// Reducer(must be default export)
export default keywordSlice.reducer;