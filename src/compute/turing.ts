
import { TuringMachine, Quadruple, Result } from 'grumle-turing';

// Create the quadruples
const quadruples: Quadruple[] = [];
quadruples.push(new Quadruple(1, 'B', 'R', 2));
quadruples.push(new Quadruple(2, '1', 'R', 2));
// etc...

// Expects an array of inputs
const inputs: number[] = [3, 2];

// Instantiate the Turing Machine
const turing: TuringMachine = new TuringMachine(quadruples, inputs);

// Use step() to walk through the execution
// const result: Result = turing.step();

/* OR */

// Use start() to execute the entire program
const result: Result = turing.start();

// Access the machine's output in decimal format
// (Only accessible once the entire machine has halted)
console.log(result.output);

// Access the tape's configuration after every step()
// or once the machine halts
console.log(result.tape);

// Access the machine's current state after every step()
// or once the machine halts
console.log(result.state);