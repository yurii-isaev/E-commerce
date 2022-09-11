import { Button, Form } from 'react-bootstrap'

const LoginPage = () => {
  return (
     <Form>
       <Form.Group className="mb-3" controlId="email">
         <Form.Label>Email address</Form.Label>
         <Form.Control type="email" placeholder="Enter email"/>
       </Form.Group>

       <Form.Group className="mb-3" controlId="password">
         <Form.Label>Password</Form.Label>
         <Form.Control type="password" placeholder="Enter password"/>
       </Form.Group>

       <Button variant="primary" type="submit">
         Submit
       </Button>
     </Form>
  )
}

export default LoginPage