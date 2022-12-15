import React from "react";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import Input from "./form/Input";

const Keyword = (props) => {
    const [keyword, setKeyword] = useState({});
    const [error, setError] = useState(null);
    const [errors, setErrors] = useState([]);
    let { id } = useParams();

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
          [name]: value,
        });
      };

    const handleSubmit = (event) => {
        event.preventDefault();
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
            name={"word"}
            value={keyword.Word}
            onChange={handleChange("keyword")}
            errorDiv={hasError("word") ? "text-danger" : "d-none"}
            errorMsg={"Please enter a keyword"}
          />
            <Input
            title={"Description"}
            className={"form-control"}
            type={"text"}
            name={"description"}
            value={keyword.Description}
            onChange={handleChange("description")}
            // errorDiv={hasError("description") ? "text-danger" : "d-none"}
            // errorMsg={"Please enter a keyword"}
          />
            <Input
            title={"ImageUrl"}
            className={"form-control"}
            type={"text"}
            name={"image_url"}
            value={keyword.ImageUrl}
            onChange={handleChange("image_url")}
            // errorDiv={hasError("description") ? "text-danger" : "d-none"}
            // errorMsg={"Please enter a keyword"}
          />
<button className="btn btn-primary">Save</button>

            {keyword.imageUrl !== "" &&
                <div className="mb-3">
                    <img src={`${keyword.ImageUrl}`} alt="picture" />
                </div>
            }

            {/* <p>{movie.description}</p> */}
            </form>
        </div>
    )
}

export default Keyword;