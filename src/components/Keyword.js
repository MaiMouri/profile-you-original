import userEvent from "@testing-library/user-event";
import React from "react";
import { useEffect, useState } from "react";
import { useSelector, useDispatch } from "react-redux";
import { useNavigate, useParams, Link } from "react-router-dom";
import Input from "./form/Input";
import {  selectKeyword, renewalKeyword } from '../keywordsSlice';



const Keyword = (props) => {
  const navigate = useNavigate();
  const [keyword, setKeyword] = useState({
    Id: 0,
    Word: "",
    Description: "",
    ImageUrl: "",
    KeywordId: ""
  });
  const [error, setError] = useState(null);
  const [errors, setErrors] = useState([]);
  let { id } = useParams();
  const { items } = useSelector(selectKeyword);
  const dispatch = useDispatch();
  
    useEffect(() => {
        const headers = new Headers();
        headers.append("Content-Type", "application/json");

        const requestOptions = {
            method: "GET",
            headers: headers,
        }

        fetch(`/keywords/${id}`, requestOptions)
            .then((response) => response.json())
            .then((data) => {
                setKeyword(data);
            })
            .catch(err => {
                console.log(err);
            })
    }, [id])

    const hasError = (key) => {
        return errors.indexOf(key) !== -1;
      };
    const handleChange = () => (event) => {
        let value = event.target.value;
        let name = event.target.name;
        setKeyword({
          ...keyword,
          [name]: value
        });
      };

    const handleSubmit = (event) => {
        event.preventDefault();

        dispatch(renewalKeyword(keyword));
        // const requestBody = keyword;
    
        // // passed validation, so save changes
        // let headers = new Headers();
        // headers.append("Content-Type", "application/json");
    
        // const requestOptions = {
        //   method: "POST",
        //   headers: headers,
        //   body: JSON.stringify(requestBody)
        // };

        // fetch(`/keyword/update/`, requestOptions)
        // .then((response) => response.json())
        // .then((data) => {
        //   if (data.error) {
        //     console.log(data.error);
        //   } else {
        //     console.log("UPDATED");
        //     navigate("/keywords");
        //   }
        // })
        // .catch((err) => {
        //   console.log(err);
        // });
      }

    return(
        <div>
            <h2>Keyword: {keyword.Word}</h2>
            <small><em>{keyword.Description}</em></small><br />
            <hr />
            <form onSubmit={handleSubmit}>
            <Input
            title={"Word"}
            className={"form-control"}
            type={"text"}
            name={"Word"}
            value={keyword.Word}
            onChange={handleChange("Keyword")}
            errorDiv={hasError("word") ? "text-danger" : "d-none"}
            errorMsg={"Please enter a keyword"}
          />
            <Input
            title={"Description"}
            className={"form-control"}
            type={"text"}
            name={"Description"}
            value={keyword.Description}
            onChange={handleChange("Description")}
            // errorDiv={hasError("description") ? "text-danger" : "d-none"}
            // errorMsg={"Please enter a keyword"}
          />
            <Input
            title={"ImageUrl"}
            className={"form-control"}
            type={"text"}
            name={"ImageUrl"}
            value={keyword.ImageUrl}
            onChange={handleChange("ImageUrl")}
            // errorDiv={hasError("description") ? "text-danger" : "d-none"}
            // errorMsg={"Please enter a keyword"}
          />
          <hr />
          <hr />
          <button className="btn btn-primary">Save</button>

            {keyword.imageUrl !== "" &&
                <div className="mb-3">
                    <img src={`${keyword.ImageUrl}`} alt="picture" />
                </div>
            }

            </form>
            <Link to={`/keywords`}><button className="btn btn-light">Home</button></Link>
        </div>
    )
}

export default Keyword;