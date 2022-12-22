// import {
//   ChakraProvider,
//   Heading,
//   Container,
//   Text,
//   Link,
//   Wrap,
//   Input,
//   Stack,
//   Button,
//   Image,
//   SkeletonCircle,
//   SkeletonText

// }
//   from "@chakra-ui/react"
// import axios from "axios";
// import Swal from "sweetalert2";
// import { useEffect, useState } from "react";
// import Keywords from "./components/Keywords";
// import React from "react";
// import { Routes, Route, useNavigate } from "react-router-dom";
// import Keyword from "./components/Keyword";
// import Login from "./components/Login";



// const App = () => {
//   const navigate = useNavigate();
//   const [image, updateImage] = useState();
//   const [word, updateWord] = useState();
//   const [loading, updateLoading] = useState();
//   const [keywords, setKeywords] = useState([]);
//   const [keyword, setKeyword] = useState({
//     Id: 0,
//     Word: "",
//     Description: "",
//     ImageUrl: "",
//     KeywordId: ""
//   });

//   const [jwtToken, setJwtToken] = useState("");
//   const [alertMessage, setAlertMessage] = useState("");
//   const [alertClassName, setAlertClassName] = useState("d-none");

//   // START FETCHING
//   useEffect(() => {
//     const headers = new Headers();
//     headers.append("Content-Type", "application/json");

//     const requestOptions = {
//       method: "GET",
//       headers: headers,
//     }

//     fetch(`http://localhost:8080/keywords`, requestOptions)
//       .then((response) => response.json())
//       .then((data) => {
//         console.log(data);
//         setKeywords(data);
//       })
//       .catch(err => {
//         console.log(err);
//       })
//   }, []);
//   // FINISH FETCHING

//   // DELETE
//   function confirmDelete(id) {

//     Swal.fire({
//       title: 'Delete keyword?',
//       text: "You cannot undo this action!",
//       icon: 'warning',
//       showCancelButton: true,
//       confirmButtonColor: '#3085d6',
//       cancelButtonColor: '#d33',
//       confirmButtonText: 'Yes, delete it!'
//     }).then((result) => {
//       if (result.isConfirmed) {
//         // axios.delete(url)
//         // .then(res => {
//         //   const keywords = this.state.keywords.filter(keyword => keyword.id !== id);
//         //   console.log("Delete from react:");
//         //   setKeywords({keywords})
//         //   console.log(res.data);
//         // })
//         let headers = new Headers();
//         headers.append("Content-Type", "application/json");
//         // headers.append("Authorization", "Bearer " + jwtToken)

//         const requestBody = keyword;
//         requestBody.KeywordId = id;

//         const requestOptions = {
//           method: "POST",
//           headers: headers,
//           body: JSON.stringify(requestBody)
//         };
        
//         const url = `http://localhost:8080/keyword/delete/`;
//         console.log(requestOptions);
//         fetch(url, requestOptions)
//           .then((response) => response.json())
//           .then((data) => {
//             if (data.error) {
//               console.log(data.error);
//             } else {
//               setKeywords((keywords) => keywords.filter((keyword) => keyword.KeywordId !== data.data));
//             }
//           })
//           .catch(err => {
//             console.log(err);
//           });
//       }
//     });
//   }
//     // FINISH DELETE

//     const handleSubmit = (e) => {
//       e.preventDefault();
//       console.log({ word });
//       generate(word)
//     };

//   const generate = async (word) => {
//     updateLoading(true);
//     // const request = await axios.post(`http://localhost:8080/keyword/create/${word}`);
//     // const result = await axios.get(`http://localhost:8080/keyword/create/${word}`);
//     // updateImage(result.data);
//     const requestBody = keyword;
//     requestBody.word = word;
    
//     let headers = new Headers();
//     headers.append("Content-Type", "application/json");
//         // headers.append("Authorization", "Bearer " + jwtToken)
//     const requestOptions = {
//       method: "POST",
//       headers: headers,
//       body: JSON.stringify(requestBody)
//     };

//         const url = `http://localhost:8080/keyword/create/${word}`;
//         fetch(url, requestOptions)
//           .then((response) => response.json())
//           .then((data) => {
//             if (data.error) {
//               console.log(data.error);
//             } else {
//               console.log(data.data)
//               setKeywords([...keywords,
//                 data.data,
//               ]);
//               // tentative first aid
//               window.location.reload();
//               navigate("/keywords")
//             }
//           })
//     updateLoading(false);
//   };


//   return (
//     <ChakraProvider>
//       <Container>
//         <Heading className="h1">Profile You🚀</Heading>
//         <Text marginBottom={"10px"}>
//           This application examines the trend of the given word in Twitter to generate images
//           using the Dall・E API. More information can be found here{" "}
//           <Link href={"#"}>
//             Web
//           </Link>
//         </Text>
//         <div className="App">
//           Profile You!
//         </div>
//         <Wrap marginBottom={"10px"}>
//           <form method="post" onSubmit={handleSubmit}>
//             <Input
//               id="word"
//               value={word}
//               name="word"
//               onChange={(e) => updateWord(e.target.value)}
//               width={"350px"}
//             ></Input>
//             <Button type="submit" colorScheme={"yellow"}>
//               Generate
//             </Button>
//           </form>
//         </Wrap>

//         {loading ? (
//           <Stack>
//             <SkeletonCircle />
//             <SkeletonText />
//           </Stack>
//         ) : image ? (
//           <Image src={`data:image/png;base64,${image}`} boxShadow="lg" />
//         ) : null}

//         <pre>{JSON.stringify(word)}</pre>

//         {/* Outletは共通NavBarとかを望むとき */}
//         {/* <Outlet context={{keywords, confirmDelete}}/> */}
//         {/* Routeの一部にしない下記の記述は居座るから🆖 */}
//         {/* <Keywords keywords={keywords} confirmDelete={confirmDelete}/> */}
//         {/* このルーティングめちゃくちゃ苦労した　何だこれ */}
//         <Routes>
//           <Route path={`/login`} element={<Login />} context={{
//               jwtToken,
//               setJwtToken,
//               setAlertClassName,
//               setAlertMessage,
//               // toggleRefresh,
//             }}/>
//           <Route path={`/keywords`} element={<Keywords keywords={keywords} confirmDelete={confirmDelete}/>} />
//           <Route path={`/keywords/:id`} element={<Keyword />} />
//         </Routes>

//       </Container>
//     </ChakraProvider>
//   );
// }

// export default App;

