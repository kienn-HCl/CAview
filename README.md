# CAview

1次元セルオートマトンをターミナル上に描画するプログラム。

## 使い方

```shell
❯ ./CAview --rule 30
                                       #                                      
                                      ###                                     
                                     #  ##                                    
                                    #### ##                                   
                                   #   #  ##                                  
                                  ### #### ##                                 
                                 #  #    #  ##                                
                                ######  #### ##                               
                               #     ###   #  ##                              
                              ###   #  ## #### ##                             
                             #  ## #### #    #  ##                            
                            #### #    # ##  #### ##                           
                           #   # ##  ##  ###   #  ##                          
                          ### ##  ### ###  ## #### ##                         
                         #  #  ###  #   ### #    #  ##                        
                        #######  ##### #  # ##  #### ##                       
                       #      ###    # ####  ###   #  ##                      
                      ###    #  ##  ##    ###  ## #### ##                     
                     #  ##  #### ### ##  #  ### #    #  ##                    
                    #### ###   #   #  ######  # ##  #### ##                   
                   #   #   ## ### ####     ####  ###   #  ##                  
                  ### ### # #   #    ##   #   ###  ## #### ##                 
                 #  #   # # ## ###  # ## ### #  ### #    #  ##                
                ###### ## #  #   ####  #   # ###  # ##  #### ##               
               #     #  # ##### #   ##### ##   ####  ###   #  ##              
              ###   #####     # ## #    #  ## #   ###  ## #### ##             
             #  ## #    ##   ##  # ##  #### # ## #  ### #    #  ##            
            #### # ##  # ## # ####  ###   # #  # ###  # ##  #### ##           
           #   # #  ####  # #    ###  ## ## ####   ####  ###   #  ##          
          ### ## ###   #### ##  #  ### #  #    ## #   ###  ## #### ##         
         #  #  #   ## #   #  ######  # #####  # # ## #  ### #    #  ##        
        ######### # # ## ####     ####     #### #  # ###  # ##  #### ##       
       #        # # #  #    ##   #   ##   #   # ####   ####  ###   #  ##      
      ###      ## # #####  # ## ### # ## ### ##    ## #   ###  ## #### ##     
     #  ##    # # #     ####  #   # #  #   #  ##  # # ## #  ### #    #  ##    
    #### ##  ## # ##   #   ##### ## ##### #### #### #  # ###  # ##  #### ##   
   #   #  ### # #  ## ### #    #  #     #    #    # ####   ####  ###   #  ##  
  ### ####  # # ### #   # ##  ######   ###  ###  ##    ## #   ###  ## #### ## 
 #  #    #### #   # ## ##  ###     ## #  ###  ### ##  # # ## #  ### #    #  ##
 #####  #   # ## ##  #  ###  ##   # # ###  ###  #  #### #  # ###  # ##  #### #
```

```shell
$ ./CAview --rule 60
                                       #                                      
                                      ##                                      
                                     # #                                      
                                    ####                                      
                                   #   #                                      
                                  ##  ##                                      
                                 # # # #                                      
                                ########                                      
                               #       #                                      
                              ##      ##                                      
                             # #     # #                                      
                            ####    ####                                      
                           #   #   #   #                                      
                          ##  ##  ##  ##                                      
                         # # # # # # # #                                      
                        ################                                      
                       #               #                                      
                      ##              ##                                      
                     # #             # #                                      
                    ####            ####                                      
                   #   #           #   #                                      
                  ##  ##          ##  ##                                      
                 # # # #         # # # #                                      
                ########        ########                                      
               #       #       #       #                                      
              ##      ##      ##      ##                                      
             # #     # #     # #     # #                                      
            ####    ####    ####    ####                                      
           #   #   #   #   #   #   #   #                                      
          ##  ##  ##  ##  ##  ##  ##  ##                                      
         # # # # # # # # # # # # # # # #                                      
        ################################                                      
       #                               #                                      
      ##                              ##                                      
     # #                             # #                                      
    ####                            ####                                      
   #   #                           #   #                                      
  ##  ##                          ##  ##                                      
 # # # #                         # # # #                                      
########                        ########                                      
```

```shell
$ ./CAview --rule 90
                                       #                                      
                                      # #                                     
                                     #   #                                    
                                    # # # #                                   
                                   #       #                                  
                                  # #     # #                                 
                                 #   #   #   #                                
                                # # # # # # # #                               
                               #               #                              
                              # #             # #                             
                             #   #           #   #                            
                            # # # #         # # # #                           
                           #       #       #       #                          
                          # #     # #     # #     # #                         
                         #   #   #   #   #   #   #   #                        
                        # # # # # # # # # # # # # # # #                       
                       #                               #                      
                      # #                             # #                     
                     #   #                           #   #                    
                    # # # #                         # # # #                   
                   #       #                       #       #                  
                  # #     # #                     # #     # #                 
                 #   #   #   #                   #   #   #   #                
                # # # # # # # #                 # # # # # # # #               
               #               #               #               #              
              # #             # #             # #             # #             
             #   #           #   #           #   #           #   #            
            # # # #         # # # #         # # # #         # # # #           
           #       #       #       #       #       #       #       #          
          # #     # #     # #     # #     # #     # #     # #     # #         
         #   #   #   #   #   #   #   #   #   #   #   #   #   #   #   #        
        # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # #       
       #                                                               #      
      # #                                                             # #     
     #   #                                                           #   #    
    # # # #                                                         # # # #   
   #       #                                                       #       #  
  # #     # #                                                     # #     # # 
 #   #   #   #                                                   #   #   #   #
  # # # # # # #                                                 # # # # # # # 
```
