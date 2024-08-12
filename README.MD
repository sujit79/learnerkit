## Learner Kit
A program which creates necessary infrastructure both mathematically and compute to enhance <br/>
ML experiments

## What it achieves
It creates an environment for creation of ML based Neural Network Experiments, currently only mathematical <br/>
Limitations are primarily due to the comparisions in golang for object, below <br/>
https://golangbyexample.com/struct-equality-golang/ <br/>

These limitations can be overcome by implementing an Equals Method <br/>
https://go.dev/ref/spec The golang documentation<br/>

## Dependencies
Dependent on https://github.com/gorgonia/gorgonia/ <br/>

## Experiment
Multiple Models are allowed and a gradient comparision of all models are done, by gradient we mean
Model 1 and Model 2 are run for example for Model 1 to Model 2. This means type safety is needed and
run time exceptions are possible, we provide error for each of these runs too. 
<br/>
The Idea is Model 1 is also run with Model 2 inputs and vice versa, one of the reasons is Model Compatibility
and a way to understand results in both Models.